package service

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo repository.PostgresUserRepository
}

func NewUserService(repo repository.PostgresUserRepository) UserService {
	return UserService{repo: repo}
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

	err = u.repo.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return dto.MakeUserOutput(user), nil
}


func (u *UserService) Update(ctx *gin.Context, id uint, user schemas.User) {

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
