package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/ExtraProjects860/Project-Device-Mobile/service"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

// @Summary      Create Enterprise
// @Description  Creates a new enterprise
// @Tags         enterprises
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        enterprise body request.EnterpriseRequest true "Enterprise info"
// @Success      201 {object} dto.EnterpriseDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/enterprise [post]
func CreateEnterprise(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input request.EnterpriseRequest

		if err := request.ReadBody(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		if err := request.ValidateBodyReq(&input, ctx, utils.GetValidate()); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		enterpriseService := service.GetEnterpriseService(appCtx)

		enterprise, err := enterpriseService.Create(ctx, input)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to create user"))
			return
		}

		response.SendSuccess(ctx, http.StatusCreated, enterprise)
	}
}

// @Summary      Update Enterprise
// @Description  Updates an existing enterprise by ID
// @Tags         enterprises
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   query int true "Enterprise ID"
// @Param        enterprise body request.EnterpriseRequest true "Enterprise info to update"
// @Success      200 {object} dto.EnterpriseDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/enterprise [put]
func UpdateEnterprise(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := request.GetIdQuery(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		var input request.EnterpriseRequest
		if err := request.ReadBody(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		if err := request.ValidateUpdateBodyReq(&input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		enterpriseService := service.GetEnterpriseService(appCtx)

		enterprise, err := enterpriseService.Update(ctx, id, input)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to update user"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, enterprise)
	}
}

// @Summary      List Enterprises
// @Description  Retrieves a paginated list of enterprises
// @Tags         enterprises
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        itemsPerPage query int true "Items per page"
// @Param        currentPage  query int true "Current page number"
// @Param        searchFilter query string false "Search item by filter"
// @Param        itemsOrder   query string false "Order direction" Enums(ASC, DESC)
// @Success      200 {array}  dto.EnterpriseDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/enterprises [get]
func GetEnterprises(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paginationSearch, err := request.GetPaginationData(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		enterpriseService := service.GetEnterpriseService(appCtx)

		enterprises, err := enterpriseService.GetAll(ctx, paginationSearch)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to process get enterprises"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, enterprises)
	}
}
