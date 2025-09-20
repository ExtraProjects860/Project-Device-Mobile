package handler

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// @Summary      Create User
// @Description  Creates a new user
// @Tags         users
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /users [post]
func CreateUser(ctx *gin.Context) {
	sendSuccess(ctx, "Create User!")
}

// @Summary      Get User Info
// @Description  Returns information about a specific user
// @Tags         users
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /users/{id} [get]
func GetInfoUser(ctx *gin.Context) {
	sendSuccess(ctx, "Get Info User!")
}

// @Summary      Get Users
// @Description  Returns a list of all users
// @Tags         users
// @Produce      json
// @Success      200 {array} map[string]string
// @Router       /users [get]
func GetUsers(ctx *gin.Context) {
	sendSuccess(ctx, "Get Users!")
}

// @Summary      Update User
// @Description  Updates an existing user
// @Tags         users
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /users/{id} [patch]
func UpdateUser(ctx *gin.Context) {
	sendSuccess(ctx, "Updated User!")
}
