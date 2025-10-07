package service

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo: repo}
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
