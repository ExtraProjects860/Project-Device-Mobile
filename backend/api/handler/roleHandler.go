package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/ExtraProjects860/Project-Device-Mobile/service"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

// @Summary      Create Role
// @Description  Creates a new role
// @Tags         roles
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        role body request.RoleRequest true "Role info"
// @Success      201 {object} dto.RoleDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/role [post]
func CreateRole(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input request.RoleRequest

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

		roleService := service.GetRoleService(appCtx)

		role, err := roleService.Create(ctx, input)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to create user"))
			return
		}

		response.SendSuccess(ctx, http.StatusCreated, role)
	}
}

// @Summary      Update Role
// @Description  Updates an existing role by ID
// @Tags         roles
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   query int true "role ID"
// @Param        role body request.RoleRequest true "Role info to update"
// @Success      200 {object} dto.RoleDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/role [put]
func UpdateRole(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := request.GetIdQuery(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		var input request.RoleRequest
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

		roleService := service.GetRoleService(appCtx)

		role, err := roleService.Update(ctx, id, input)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to update user"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, role)
	}
}

// @Summary      List Roles
// @Description  Retrieves a paginated list of roles
// @Tags         roles
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        itemsPerPage query int true "Items per page"
// @Param        currentPage  query int true "Current page number"
// @Success      200 {array}  dto.RoleDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/roles [get]
func GetRoles(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		itemsPerPage, currentPage, err := request.GetPaginationData(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		roleService := service.GetRoleService(appCtx)

		roles, err := roleService.GetAll(ctx, itemsPerPage, currentPage)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to process get roles"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, roles)
	}
}
