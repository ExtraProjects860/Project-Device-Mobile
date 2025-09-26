package handler

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary      Create Product
// @Description  Creates a new product
// @Tags         products
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /product [post]
func CreateProductHandler(ctx *gin.Context, repo repository.ProductRepository) {
	sendSuccess(ctx, gin.H{"message": "Add Promotion Product!"})
}

// @Summary      Update Product
// @Description  Updates an existing product
// @Tags         products
// @Param 		 id query string true "Product ID"
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /product [patch]
func UpdateProductHandler(ctx *gin.Context, repo repository.ProductRepository) {
	sendSuccess(ctx, gin.H{"message": "Update Promotion Product!"})
}

// @Summary      Get Products
// @Description  Returns all products
// @Tags         products
// @Produce      json
// @Success      200 {array} map[string]string
// @Router       /products [get]
func GetProductsHandler(ctx *gin.Context, repo repository.ProductRepository) {
	sendSuccess(ctx, gin.H{"message": "Get Promotions Products!"})
}
