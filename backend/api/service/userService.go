package service

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo   *repository.PostgresUserRepository
	logger *config.Logger
}

func GetUserService(appCtx *appcontext.AppContext) UserService {
	return UserService{
		repo:   repository.NewPostgresUserRepository(appCtx.DB),
		logger: config.NewLogger("SERVICE - USER"),
	}
}

func (u *UserService) ValidateAndUpdateFields(user *schemas.User, input request.UserRequest) error {
	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		hashed, err := utils.GenerateHashPassword(input.Password)
		if err != nil {
			return fmt.Errorf("password hash: %v", err)
		}
		user.Password = hashed
	}
	if input.Cpf != "" && utils.ValidateCPF(input.Cpf) {
		user.Cpf = input.Cpf
	}
	if input.RegisterNumber != "" && len(input.RegisterNumber) == 7 {
		user.RegisterNumber = input.RegisterNumber
	}
	if input.RoleID != 0 {
		user.RoleID = input.RoleID
	}
	if input.EnterpriseID != nil && *input.EnterpriseID != 0 {

		user.EnterpriseID = input.EnterpriseID
	}
	if input.PhotoUrl != nil && *input.PhotoUrl != "" {
		user.PhotoUrl = input.PhotoUrl
	}
	return nil
}

func (u *UserService) Create(ctx *gin.Context, input request.UserRequest) (*dto.UserDTO, error) {
	hashedPassword, err := utils.GenerateHashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := schemas.User{
		RoleID:         input.RoleID,
		EnterpriseID:   input.EnterpriseID,
		Name:           input.Name,
		Email:          input.Email,
		Password:       hashedPassword,
		Cpf:            input.Cpf,
		RegisterNumber: input.RegisterNumber,
		PhotoUrl:       input.PhotoUrl,
	}

	if err = u.repo.CreateUser(ctx, &user); err != nil {
		return nil, err
	}

	return dto.MakeUserOutput(user), nil
}

func (u *UserService) Update(ctx *gin.Context, id uint, input request.UserRequest) (*dto.UserDTO, error) {
	user, err := u.repo.GetInfoUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if err = u.ValidateAndUpdateFields(&user, input); err != nil {
		return nil, err
	}

	if err = u.repo.UpdateUser(ctx, id, &user); err != nil {
		return nil, err
	}

	return dto.MakeUserOutput(user), nil
}

func (u *UserService) Get(ctx *gin.Context, id uint) (*dto.UserDTO, error) {
	user, err := u.repo.GetInfoUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.MakeUserOutput(user), nil
}

func (u *UserService) GetAll(ctx *gin.Context, paginationSearch request.PaginationSearch) (*dto.PaginationDTO, error) {
	users, totalPages, totalItems, err := u.repo.GetUsers(
		ctx,
		paginationSearch,
	)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}

	toDTO := func(user schemas.User) *dto.UserDTO {
		return dto.MakeUserOutput(user)
	}

	return dto.MakePaginationDTO(
		users,
		paginationSearch.CurrentPage,
		totalPages,
		totalItems,
		toDTO,
	)
}
