package handler

import (
	"net/http"

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
// @Param        itemsPerPage query string true "Pagination Items"
// @Param        currentPage query string true "Pagination Current Page"
// @Success      200 {array}  repository.ProductDTO
// @Failure      400 {object} ErrResponse
// @Failure      500 {object} ErrResponse
// @Router       /products [get]
func GetProductsHandler(ctx *gin.Context, repo repository.ProductRepository) {
	itemsPerPage, currentPage, err := getPaginationData(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	products, err := repo.GetProducts(ctx, itemsPerPage, currentPage)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, gin.H{"error": "Error to get products in database"})
		return
	}

	sendSuccess(ctx, products)
}
