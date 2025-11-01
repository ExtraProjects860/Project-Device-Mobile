package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

/*
TODO essa parte do código em especial merece uma atenção posteriormente
OBS: pode ser que seja util utilizar por exemplo TokenPassword.User e acessar as coisas lá
apenas fazendo um Preload e provavelmente eliminando de ter que usar o usuário
*/

/*
TODO revisar o uso de logs dentro dos services, pois alguns são desnecessários outros inúteis
avaliar também onde pode ser info, debug, warning, error e critical
*/

type AuthService struct {
	logger      *config.Logger
	repo        repository.PostgresAuthRepository
	userService UserService
}

type PayloadEmail struct {
	Subject      string         `json:"subject"`
	SendTo       string         `json:"send_to"`
	TemplateName string         `json:"template_name"`
	Data         map[string]any `json:"data"`
}

func GetAuthService(appCtx *appcontext.AppContext) AuthService {
	return AuthService{
		logger:      config.NewLogger("SERVICE - AUTH"),
		repo:        *repository.NewPostgresAuthRepository(appCtx.DB),
		userService: GetUserService(appCtx),
	}
}

func (s *AuthService) generateTempAndCode() (string, time.Time) {
	code := utils.GenerateRandomCode(6)
	timeUp := time.Now().Add(1 * time.Minute)

	return code, timeUp
}

func (s *AuthService) sendMailPassword(ctx *gin.Context, apiURI string, payload PayloadEmail) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		s.logger.Errorf("Error marshaling JSON: %v", err)
		return errors.New("error marshaling JSON")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		apiURI+"/email",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		s.logger.Errorf("Error creating request: %v", err)
		return errors.New("error creating request")
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Errorf("Error sending request for email %v", err)
		return errors.New("error sending request for email")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.logger.Errorf("error to send email for: %v", payload.SendTo)
		return fmt.Errorf("error to send email for: %v", payload.SendTo)
	}

	return nil
}

func (s *AuthService) verifyTimesUpToken(ctx *gin.Context, tokenSchema *schemas.TokenPassword) bool {
	return !(tokenSchema.TimeUp == nil || time.Now().After(*tokenSchema.TimeUp))
}

func (s *AuthService) GetToken(ctx *gin.Context, id uint) (schemas.TokenPassword, error) {
	tokenSchema, err := s.repo.GetToken(ctx, id)
	if err != nil {
		s.logger.Errorf("Error retrieving existing token for UserID %d: %v", id, err)
		return schemas.TokenPassword{}, errors.New("error verifying existing token")
	}

	return tokenSchema, nil
}

func (s *AuthService) CreateToken(ctx *gin.Context, apiURI string, user *dto.UserDTO) error {
	tokenSchema, err := s.GetToken(ctx, user.ID)
	if err != nil {
		return err
	}

	code, timeUp := s.generateTempAndCode()

	if tokenSchema.ID == 0 {

		newToken := &schemas.TokenPassword{
			UserID: user.ID,
			Code:   &code,
			TimeUp: &timeUp,
		}

		if err := s.repo.CreateToken(ctx, newToken); err != nil {
			s.logger.Errorf("Error creating new token for UserID %d: %v", user.ID, err)
			return errors.New("error saving new token")
		}
	} else {
		isUp := s.verifyTimesUpToken(ctx, &tokenSchema)
		if isUp {
			return errors.New("token already exists and is still valid, wait until it expires")
		}

		tokenSchema.Code = &code
		tokenSchema.TimeUp = &timeUp

		if err := s.repo.UpdateToken(ctx, tokenSchema.ID, &tokenSchema); err != nil {
			s.logger.Errorf("Error updating token for UserID %d: %v", user.ID, err)
			return errors.New("error saving updated token")
		}
	}

	err = s.sendMailPassword(ctx, apiURI, PayloadEmail{
		Subject:      "Token para troca de senha",
		SendTo:       user.Email,
		TemplateName: "change_password",
		Data:         map[string]any{"token": code},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ChangePassword(
	ctx *gin.Context,
	apiURI string,
	code string,
	newPassword string,
	user *dto.UserDTO,
	repo *repository.PostgresUserRepository,
) error {
	tokenSchema, err := s.GetToken(ctx, user.ID)
	if err != nil {
		return err
	}

	isUp := s.verifyTimesUpToken(ctx, &tokenSchema)
	if !isUp {
		return errors.New("token is expired")
	}
	if *tokenSchema.Code != code {
		s.logger.Warningf("Invalid token code provided")
		return errors.New("invalid token code provided")
	}

	err = s.ResetPassword(ctx, newPassword, user.ID, repo)
	if err != nil {
		return err
	}

	//TODO posso mover isso para um serviço de email separado depois para melhorar o código
	err = s.sendMailPassword(ctx, apiURI, PayloadEmail{
		Subject:      "Senha Alterada com Sucesso",
		SendTo:       user.Email,
		TemplateName: "confirmation",
		Data:         map[string]any{"token": code},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) VerifyCredentials(
	ctx *gin.Context,
	input request.LoginRequest,
	repo *repository.PostgresUserRepository,
) (uint, error) {
	user, err := repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		s.logger.Warning(err.Error())
		return 0, fmt.Errorf("error incorrect email or not found, try again or other")
	}

	err = utils.VerifyHashedPassword(input.Password, user.Password)
	if err != nil {
		s.logger.Warning(err.Error())
		return 0, fmt.Errorf("password incorrect, try again or other")
	}

	return user.ID, nil
}

func (s *AuthService) ResetPassword(
	ctx *gin.Context,
	newPassword string,
	id uint,
	repo *repository.PostgresUserRepository,
) error {
	user, err := repo.GetInfoUser(ctx, id)
	if err != nil {
		s.logger.Warning(err.Error())
		return fmt.Errorf("error not found use by id")
	}

	hashedPassword, err := utils.GenerateHashPassword(newPassword)
	if err != nil {
		s.logger.Warning(err.Error())
		return fmt.Errorf("error to generate password hash")
	}

	user.Password = hashedPassword

	if err := repo.UpdateUser(ctx, id, &user); err != nil {
		s.logger.Warning(err.Error())
		return fmt.Errorf("error to update password user")
	}

	return nil
}
