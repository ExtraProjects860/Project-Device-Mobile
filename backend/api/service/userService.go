package service

import (
	"context"
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
	"github.com/paemuri/brdoc"
)

type UserService struct {
	repo repository.PostgresUserRepository
}

func NewUserService(repo repository.PostgresUserRepository) UserService {
	return UserService{repo: repo}
}

func validateAndUpdateFields(user *schemas.User, input request.UserRequest) error {
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
	if input.Cpf != "" && brdoc.IsCPF(input.Cpf) {
		user.Cpf = input.Cpf
	}
	if input.RegisterNumber != 0 {
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

func (u *UserService) Create(ctx context.Context, input request.UserRequest) (*dto.UserDTO, error) {
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

	if err = validateAndUpdateFields(&user, input); err != nil {
		return nil, err
	}

	if err = u.repo.UpdateUser(ctx, id, &user); err != nil {
		return nil, err
	}

	return dto.MakeUserOutput(user), nil
}

func (u *UserService) Get(ctx context.Context, id uint) (*dto.UserDTO, error) {
	user, err := u.repo.GetInfoUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.MakeUserOutput(user), nil
}

func (u *UserService) GetAll(ctx *gin.Context, itemsPerPage, currentPage uint) (*dto.PaginationDTO, error) {
	users, totalPages, totalItems, err := u.repo.GetUsers(ctx, itemsPerPage, currentPage)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	toDTO := func(user schemas.User) *dto.UserDTO {
		return dto.MakeUserOutput(user)
	}

	return dto.MakePaginationDTO(
		users,
		currentPage,
		totalPages,
		totalItems,
		toDTO,
	)
}
