package handler

import (
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary      Add Product to Wish List
// @Description  Adds a product to the user wish list
// @Tags         wishlist
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /wishlist [post]
func AddInWishListHandler(ctx *gin.Context, repo repository.WishListRepository) {
	sendSuccess(ctx, "Add Product in Wish List!")
}

// @Summary      Update Product from Wish List
// @Description  Update a product from the user wish list
// @Tags         wishlist
// @Param 		 id query string true "WishList ID"
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /wishlist [patch]
func UpdateWishListHandler(ctx *gin.Context, repo repository.WishListRepository) {
	sendSuccess(ctx, "Delete Product in Wish List!")
}

// @Summary      Get Wish List Items
// @Description  Returns all products in the user wish list
// @Tags         wishlist
// @Produce      json
// @Param 		 id query string true "User ID"
// @Param        itemsPerPage query string true "Pagination Items"
// @Param        currentPage query string true "Pagination Current Page"
// @Success      200 {array}  repository.WishListDTO
// @Failure      400 {object} ErrResponse
// @Failure      500 {object} ErrResponse
// @Router       /wishlist [get]
func GetWishListByUserIDHandler(ctx *gin.Context, repo repository.WishListRepository) {
	userId, err := getIdQuery(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	itemsPerPage, currentPage, err := getPaginationData(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wishlist, err := repo.GetWishListByUserID(ctx, userId, itemsPerPage, currentPage)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, gin.H{"error": "Error to get wishlist in database"})
		return
	}

	sendSuccess(ctx, wishlist)
}
