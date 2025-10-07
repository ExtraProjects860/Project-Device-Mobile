package service

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/gin-gonic/gin"
)

type WishListService struct {
	repo repository.WishListRepository
}

func NewWishListService(repo repository.WishListRepository) WishListService {
	return WishListService{repo: repo}
}

func (w *WishListService) GetAll(ctx *gin.Context, userID, itemsPerPage, currentPage uint) (*dto.PaginationDTO, error) {
	wishlists, totalPages, totalItems, err := w.repo.GetWishListByUserID(ctx, userID, itemsPerPage, currentPage)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	wishlistDTO := dto.MakeWishListOutput(wishlists, userID)

	return &dto.PaginationDTO{
		Data:        wishlistDTO,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}, nil
}
