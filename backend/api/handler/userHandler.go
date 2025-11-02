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

// @Summary      Create User
// @Description  Creates a new user
// @Tags         users
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        image formData file false "Optional user profile image"
// @Param        data formData string true "JSON string contain user data for create (request.UserRequest)"
// @Success      201 {object} dto.UserDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/user [post]
func CreateUserHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input request.UserRequest

		if err := request.ReadBodyFORM(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		if err := request.ValidateBodyReq(&input, ctx, utils.GetValidate()); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		userService := service.GetUserService(appCtx)

		user, err := userService.Create(
			ctx, service.GetImageService(appCtx), input)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, err)
			return
		}

		response.SendSuccess(ctx, http.StatusCreated, user)
	}
}

// @Summary      Get User Info
// @Description  Returns information about a specific user
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} dto.UserDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/user [get]
func GetInfoUserHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := request.GetIdByToken(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnauthorized, err)
			return 
		}

		userService := service.GetUserService(appCtx)

		user, err := userService.Get(ctx, uid)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to process get user"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, user)
	}
}

// @Summary      Get Users
// @Description  Returns a list of all users
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        itemsPerPage query string true "Pagination Items"
// @Param        currentPage query string true "Pagination Current Page"
// @Param        searchFilter query string false "Search item by filter"
// @Param        itemsOrder   query string false "Order direction" Enums(ASC, DESC)
// @Success      200 {array}  dto.UserDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/users [get]
func GetUsersHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paginationSearch, err := request.GetPaginationData(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		userService := service.GetUserService(appCtx)

		users, err := userService.GetAll(ctx, paginationSearch)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to process get users"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, users)
	}
}

// @Summary      Update User
// @Description  Updates an existing user
// @Tags         users
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param 		 id query string true "User ID"
// @Param        image formData file false "Optional user profile image"
// @Param        data formData string true "JSON string contain user data for update (request.UserRequest)"
// @Success      200 {object} dto.UserDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/user [patch]
func UpdateUserHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := request.GetIdQuery(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		var input request.UserRequest
		if err := request.ReadBodyFORM(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		if err := request.ValidateUpdateBodyReq(&input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		userService := service.GetUserService(appCtx)
		imageService := service.GetImageService(appCtx)

		user, err := userService.Update(ctx, imageService, id, input)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to update user"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, user)
	}
}
