package handler

import (
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary      Create User
// @Description  Creates a new user
// @Tags         users
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /user [post]
func CreateUserHandler(ctx *gin.Context) {
	sendSuccess(ctx, "Create User!")
}

// @Summary      Get User Info
// @Description  Returns information about a specific user
// @Tags         users
// @Param 		 id query string true "User ID"
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /user [get]
func GetInfoUserHandler(ctx *gin.Context) {
	err := getIdQuery(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sendSuccess(ctx, "Get Info User!")
}

// @Summary      Get Users
// @Description  Returns a list of all users
// @Tags         users
// @Produce      json
// @Success      200 {array} schemas.UserResponse
// @Failure      400 {object} ErrResponse
// @Failure      500 {object} ErrResponse
// @Router       /users [get]
func GetUsersHandler(ctx *gin.Context) {
	repo := repository.NewPostgresUserRepository()
	users, err := repo.GetUsers(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, gin.H{"error": "Error to get users in database"})
		return
	}
	sendSuccess(ctx, users)
}

// @Summary      Update User
// @Description  Updates an existing user
// @Tags         users
// @Param 		 id query string true "User ID"
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /user [patch]
func UpdateUserHandler(ctx *gin.Context) {
	err := getIdQuery(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sendSuccess(ctx, "Updated User!")
}
