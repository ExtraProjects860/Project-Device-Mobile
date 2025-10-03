package handler

import (
	"net/http"
<<<<<<< HEAD

	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)
=======
>>>>>>> dev

	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// @Summary      Create Product
// @Description  Creates a new product
// @Tags         products
// @Produce      json
// @Success      200 {object} map[string]string
<<<<<<< HEAD
// @Router       /product [post]
func CreateProductHandler(ctx *gin.Context, repo repository.ProductRepository) {
=======
// @Router       /api/v1/product [post]
func (h *ProductHandler) CreateProductHandler(ctx *gin.Context) {
>>>>>>> dev
	sendSuccess(ctx, gin.H{"message": "Add Promotion Product!"})
}

// @Summary      Update Product
// @Description  Updates an existing product
// @Tags         products
// @Param 		 id query string true "Product ID"
// @Produce      json
// @Success      200 {object} map[string]string
<<<<<<< HEAD
// @Router       /product [patch]
func UpdateProductHandler(ctx *gin.Context, repo repository.ProductRepository) {
=======
// @Router       /api/v1/product [patch]
func (h *ProductHandler) UpdateProductHandler(ctx *gin.Context) {
>>>>>>> dev
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
<<<<<<< HEAD
// @Router       /products [get]
func GetProductsHandler(ctx *gin.Context, repo repository.ProductRepository) {
=======
// @Router       /api/v1/products [get]
func (h *ProductHandler) GetProductsHandler(ctx *gin.Context) {
>>>>>>> dev
	itemsPerPage, currentPage, err := getPaginationData(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

<<<<<<< HEAD
	products, err := repo.GetProducts(ctx, itemsPerPage, currentPage)
=======
	products, err := h.repo.GetProducts(ctx, itemsPerPage, currentPage)
>>>>>>> dev
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, gin.H{"error": "Error to get products in database"})
		return
	}

	sendSuccess(ctx, products)
}
