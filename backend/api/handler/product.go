package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/service"
	"github.com/gin-gonic/gin"
)

// @Summary      Create Product
// @Description  Creates a new product
// @Tags         products
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /api/v1/product [post]
func CreateProductHandler(ctx *gin.Context) {
	sendSuccess(ctx, http.StatusCreated, gin.H{"message": "Add Promotion Product!"})
}

// @Summary      Update Product
// @Description  Updates an existing product
// @Tags         products
// @Param 		 id query string true "Product ID"
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /api/v1/product [patch]
func UpdateProductHandler(ctx *gin.Context) {
	sendSuccess(ctx, http.StatusOK, gin.H{"message": "Update Promotion Product!"})
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
// @Router       /api/v1/products [get]
func GetProductsHandler(ctx *gin.Context) {
	itemsPerPage, currentPage, err := getPaginationData(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusBadRequest, err)
		return
	}

	productService := service.NewProductService(
		repository.NewPostgresProductRepository(config.GetDB()),
	)

	products, err := productService.GetAll(ctx, itemsPerPage, currentPage)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, errors.New("error to get products in database"))
		return
	}

	sendSuccess(ctx, http.StatusOK, products)
}
