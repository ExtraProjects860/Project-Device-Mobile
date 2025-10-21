package service

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	logger *config.Logger
}

func GetAuthService(appCtx *appcontext.AppContext) AuthService {
	return AuthService{
		logger: config.NewLogger("SERVICE - AUTH"),
	}
}

func (s *AuthService) VerifyCredentials(ctx *gin.Context, input request.LoginRequest, repo *repository.PostgresUserRepository) (uint, error) {
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

func (s *AuthService) ResetPasswordLogIn(ctx *gin.Context, newPassword string, id uint, repo *repository.PostgresUserRepository) error {
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
