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
// @Accept       json
// @Produce      json
// @Param        user body request.UserRequest false "User info"
// @Success      201 {object} dto.UserDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/user [post]
func CreateUserHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input request.UserRequest

		if err := request.ReadBody(ctx, &input); err != nil {
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

		user, err := userService.Create(ctx, input)
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
// @Param 		 id query string true "User ID"
// @Produce      json
// @Success      200 {object} dto.UserDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/user [get]
func GetInfoUserHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := request.GetIdQuery(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		userService := service.GetUserService(appCtx)

		user, err := userService.Get(ctx, id)
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
// @Produce      json
// @Param        itemsPerPage query string true "Pagination Items"
// @Param        currentPage query string true "Pagination Current Page"
// @Success      200 {array}  dto.UserDTO
// @Failure      400 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/users [get]
func GetUsersHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		itemsPerPage, currentPage, err := request.GetPaginationData(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		userService := service.GetUserService(appCtx)

		users, err := userService.GetAll(ctx, itemsPerPage, currentPage)
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
// @Accept       json
// @Produce      json
// @Param 		 id query string true "User ID"
// @Param        user body request.UserRequest false "User info to update"
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

		userService := service.GetUserService(appCtx)

		user, err := userService.Update(ctx, id, input)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to update user"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, user)
	}
}
