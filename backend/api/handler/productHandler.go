package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/ExtraProjects860/Project-Device-Mobile/service"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

// @Summary      Create Product
// @Description  Creates a new product
// @Tags         products
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        image formData file false "Optional product profile image"
// @Param        data formData string true "JSON string contain product data for create (request.ProductRequest)"
// @Success      201 {object} dto.ProductDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/product [post]
func CreateProductHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input request.ProductRequest

		if err := request.ReadBodyFORM(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		if err := request.ValidateBodyReq(&input, ctx, utils.GetValidate()); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		productService := service.GetProductService(appCtx)
		imageService := service.GetImageService(appCtx)

		product, err := productService.Create(ctx, imageService, input)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, err)
			return
		}

		response.SendSuccess(ctx, http.StatusCreated, product)
	}
}

// @Summary      Update Product
// @Description  Updates an existing product
// @Tags         products
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param 		 id query string true "Product ID"
// @Param        image formData file false "Optional product profile image"
// @Param        data formData string true "JSON string contain product data for update (request.ProductRequest)"
// @Success      200 {object} dto.ProductDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/product [patch]
func UpdateProductHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := request.GetIdQuery(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		var input request.ProductRequest
		if err := request.ReadBodyFORM(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		fmt.Println(input)

		if err := request.ValidateUpdateBodyReq(&input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		productService := service.GetProductService(appCtx)
		imageService := service.GetImageService(appCtx)

		product, err := productService.Update(ctx, imageService, id, input)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to update product"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, product)
	}
}

// @Summary      Get Products
// @Description  Returns all products
// @Tags         products
// @Security     BearerAuth
// @Produce      json
// @Param        itemsPerPage query string true "Pagination Items"
// @Param        currentPage query string true "Pagination Current Page"
// @Param        searchFilter query string false "Search item by filter"
// @Param        itemsOrder   query string false "Order direction" Enums(ASC, DESC)
// @Success      200 {array}  dto.ProductDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/products [get]
func GetProductsHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paginationSearch, err := request.GetPaginationData(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		productService := service.GetProductService(appCtx)

		products, err := productService.GetAll(ctx, paginationSearch)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to get products in database"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, products)
	}
}
