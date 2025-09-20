package handler

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// @Summary      Create Product
// @Description  Creates a new product
// @Tags         products
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /products [post]
func CreateProduct(ctx *gin.Context) {
	sendSuccess(ctx, "Add Promotion Product!")
}

// @Summary      Update Product
// @Description  Updates an existing product
// @Tags         products
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /products/{id} [patch]
func UpdateProduct(ctx *gin.Context) {
	sendSuccess(ctx, "Update Promotion Product!")
}

// @Summary      Get Products
// @Description  Returns all products
// @Tags         products
// @Produce      json
// @Success      200 {array} map[string]string
// @Router       /products [get]
func GetProducts(ctx *gin.Context) {
	sendSuccess(ctx, "Get Promotions Products!")
}
