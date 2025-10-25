package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/ExtraProjects860/Project-Device-Mobile/service"
	"github.com/gin-gonic/gin"
)

// @Summary      Add Product to Wish List
// @Description  Adds a product to the user wish list
// @Tags         wishlist
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param 		 user_id query string true "User ID"
// @Param 		 product_id query string true "Product ID"
// @Success      201 {object} dto.WishListMinimalDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/wishlist [post]
func AddInWishListHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, productID, err := request.GetUserAndProductIdQuery(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		wishlistService := service.GetWishListService(appCtx)

		wishlistEntrie, err := wishlistService.Create(ctx, userID, productID)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, err)
			return
		}

		response.SendSuccess(ctx, http.StatusCreated, wishlistEntrie)
	}
}

// @Summary      Delete Product from Wish List
// @Description  Delete a product from the user wish list
// @Tags         wishlist
// @Security     BearerAuth
// @Produce      json
// @Param 		 user_id query string true "User ID"
// @Param 		 product_id query string true "Product ID"
// @Success      200 {object} dto.WishListMinimalDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/wishlist [delete]
func DeleteInWishListHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, productID, err := request.GetUserAndProductIdQuery(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		wishlistService := service.GetWishListService(appCtx)

		wishlistEntrie, err := wishlistService.Delete(ctx, userID, productID)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to delete item in wishlist"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, wishlistEntrie)
	}
}

// @Summary      Get Wish List Items
// @Description  Returns all products in the user wish list
// @Tags         wishlist
// @Security     BearerAuth
// @Produce      json
// @Param        itemsPerPage query string true "Pagination Items"
// @Param        currentPage query string true "Pagination Current Page"
// @Param        searchFilter query string false "Search item by filter"
// @Param        itemsOrder   query string false "Order direction" Enums(ASC, DESC)
// @Success      200 {array}  dto.ProductDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/wishlist [get]
func GetWishListByUserIDHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uidRaw, exists := ctx.Get("user_id")
		if !exists {
			response.SendErr(ctx, http.StatusUnauthorized, errors.New("user id not found in token"))
			return
		}

		uid, ok := uidRaw.(uint)
		if !ok {
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("invalid user id type"))
			return
		}

		paginationSearch, err := request.GetPaginationData(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		wishlistService := service.GetWishListService(appCtx)

		wishlistEntries, err := wishlistService.GetAll(ctx, uid, paginationSearch)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to get wishlist in database"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, wishlistEntries)
	}
}
