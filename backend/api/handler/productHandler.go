package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/service"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

// @Summary      Create Product
// @Description  Creates a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body request.ProductRequest true "Product info"
// @Success      201 {object} dto.ProductDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/product [post]
func CreateProductHandler(ctx *gin.Context) {
	var input request.ProductRequest

	if err := request.ReadBody(ctx, &input); err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusBadRequest, err)
		return
	}

	if err := request.ValidateBodyReq(&input, utils.GetValidate()); err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusBadRequest, err)
		return
	}

	productService := service.NewProductService(
		*repository.NewPostgresProductRepository(config.GetDB()),
	)

	product, err := productService.Create(ctx, input)
	if err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusInternalServerError, err)
		return
	}

	response.SendSuccess(ctx, http.StatusCreated, product)
}

// @Summary      Update Product
// @Description  Updates an existing product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param 		 id query string true "Product ID"
// @Param        product body request.ProductRequest true "Product info to update"
// @Success      200 {object} dto.ProductDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/product [patch]
func UpdateProductHandler(ctx *gin.Context) {
	id, err := request.GetIdQuery(ctx)
	if err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusBadRequest, err)
		return
	}

	var input request.ProductRequest
	if err := request.ReadBody(ctx, &input); err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusBadRequest, err)
		return
	}

	if err := request.ValidateUpdateBodyReq(&input); err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusBadRequest, err)
		return
	}

	productService := service.NewProductService(
		*repository.NewPostgresProductRepository(config.GetDB()),
	)

	user, err := productService.Update(ctx, id, input)
	if err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to update user"))
		return
	}

	response.SendSuccess(ctx, http.StatusOK, user)
}

// @Summary      Get Products
// @Description  Returns all products
// @Tags         products
// @Produce      json
// @Param        itemsPerPage query string true "Pagination Items"
// @Param        currentPage query string true "Pagination Current Page"
// @Success      200 {array}  dto.ProductDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/products [get]
func GetProductsHandler(ctx *gin.Context) {
	itemsPerPage, currentPage, err := request.GetPaginationData(ctx)
	if err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusBadRequest, err)
		return
	}

	productService := service.NewProductService(
		*repository.NewPostgresProductRepository(config.GetDB()),
	)

	products, err := productService.GetAll(ctx, itemsPerPage, currentPage)
	if err != nil {
		logger.Error(err.Error())
		response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to get products in database"))
		return
	}

	response.SendSuccess(ctx, http.StatusOK, products)
}
