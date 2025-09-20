package handler

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// @Summary      Add Product to Wish List
// @Description  Adds a product to the user wish list
// @Tags         wishlist
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /wishlist [post]
func AddInWishListHandler(ctx *gin.Context) {
	sendSuccess(ctx, "Add Product in Wish List!")
}

// @Summary      Delete Product from Wish List
// @Description  Removes a product from the user wish list
// @Tags         wishlist
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /wishlist/{id} [delete]
func UpdateWishListHandler(ctx *gin.Context) {
	sendSuccess(ctx, "Delete Product in Wish List!")
}

// @Summary      Get Wish List Items
// @Description  Returns all products in the user wish list
// @Tags         wishlist
// @Produce      json
// @Success      200 {array} map[string]string
// @Router       /wishlist [get]
func GetItensWishListHandler(ctx *gin.Context) {
	sendSuccess(ctx, "Get Itens in Wish List!")
}
