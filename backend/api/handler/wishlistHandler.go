package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/service"
	"github.com/gin-gonic/gin"
)

// @Summary      Add Product to Wish List
// @Description  Adds a product to the user wish list
// @Tags         wishlist
// @Accept       json
// @Produce      json
// @Param        wishlist body request.WishListRequest true "WishList info"
// @Success      201 {object} dto.WishListDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/wishlist [post]
func AddInWishListHandler(ctx *gin.Context) {
	response.SendSuccess(ctx, http.StatusCreated, "Add Product in Wish List!")
}

// @Summary      Delete Product from Wish List
// @Description  Delete a product from the user wish list
// @Tags         wishlist
// @Produce      json
// @Param 		 user_id query string true "User ID"
// @Param 		 product_id query string true "Product ID"
// @Success      200 {object} map[string]string
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/wishlist [delete]
func DeleteInWishListHandler(ctx *gin.Context) {
	response.SendSuccess(ctx, http.StatusOK, "Delete Product in Wish List!")
}

// @Summary      Get Wish List Items
// @Description  Returns all products in the user wish list
// @Tags         wishlist
// @Produce      json
// @Param 		 id query string true "User ID"
// @Param        itemsPerPage query string true "Pagination Items"
// @Param        currentPage query string true "Pagination Current Page"
// @Success      200 {array}  dto.WishListDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/wishlist [get]
func GetWishListByUserIDHandler(ctx *gin.Context) {
	userId, err := request.GetIdQuery(ctx)
	if err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusBadRequest, err)
		return
	}

	itemsPerPage, currentPage, err := request.GetPaginationData(ctx)
	if err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusBadRequest, err)
		return
	}

	wishlistService := service.NewWishListService(
		*repository.NewPostgresWishListRepository(config.GetDB()),
	)

	wishlistEntries, err := wishlistService.GetAll(ctx, userId, itemsPerPage, currentPage)
	if err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to get wishlist in database"))
		return
	}

	response.SendSuccess(ctx, http.StatusOK, wishlistEntries)
}
