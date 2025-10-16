package service

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/service/dto"
	"github.com/gin-gonic/gin"
)

type WishListService struct {
	repo   *repository.PostgresWishListRepository
	logger *config.Logger
}

func GetWishListService(appCtx *appcontext.AppContext) WishListService {
	return WishListService{
		repository.NewPostgresWishListRepository(appCtx.DB),
		config.NewLogger("SERVICE - WISHLIST"),
	}
}

func (w *WishListService) Create(ctx *gin.Context, userID, productID uint) (*dto.WishListMinimalDTO, error) {
	wishlist := &schemas.WishList{
		UserID:    userID,
		ProductID: productID,
	}

	if err := w.repo.AddInWishList(ctx, wishlist); err != nil {
		return nil, err
	}

	return dto.MakeWishListMinimalDTO(*wishlist), nil
}

func (w *WishListService) Delete(ctx *gin.Context, userID, productID uint) (*dto.WishListMinimalDTO, error) {
	err := w.repo.DeleteInWishList(ctx, userID, productID)
	if err != nil {
		return nil, err
	}

	return dto.MakeWishListMinimalDTO(schemas.WishList{
		UserID:    userID,
		ProductID: productID,
	}), nil
}

func (w *WishListService) GetAll(ctx *gin.Context, userID, itemsPerPage, currentPage uint) (*dto.PaginationDTO, error) {
	wishlists, totalPages, totalItems, err := w.repo.GetWishListByUserID(ctx, userID, itemsPerPage, currentPage)
	if err != nil {
		w.logger.Error(err.Error())
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
