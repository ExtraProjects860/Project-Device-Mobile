package handler

import (
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/auth"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary      Request Password Token
// @Description  Requests a reset token for user password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /auth/request-token [post]
func RequestTokenHandler(ctx *gin.Context) {
	sendSuccess(ctx, "Require Password!")
}

// @Summary      Reset Password
// @Description  Resets user password using the token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /auth/reset-password [post]
func ResetPasswordHandler(ctx *gin.Context) {
	sendSuccess(ctx, "Change Password!")
}

// @Summary      User Login
// @Description  Authenticates user and returns access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body LoginRequest true "Request body"
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /auth/login [post]
func LoginHandler(ctx *gin.Context) {
	var request LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		sendErr(ctx, http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	// TODO Simulating user authentication (replace with real logic) *se não vai dar merda, e precisa implementar o repository*
	if request.Email != "test@gmail.com" || request.Password != "1234ok" {
		sendErr(ctx, http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	user := &UserData{ID: 1, Email: "test@gmail.com"}
	jwtToken, refreshToken, err := auth.GenerateTokens(user.ID)
	if err != nil {
		sendErr(ctx, http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	sendSuccess(ctx, gin.H{
		"jwt_token":     jwtToken,
		"refresh_token": refreshToken,
	})
}

// @Summary      Refresh Token
// @Description  Refreshes the authentication token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /auth/refresh-token [post]
func RefreshTokenHandler(ctx *gin.Context) {
	var request auth.RequestRefresh
	if err := ctx.ShouldBindJSON(&request); err != nil {
		sendErr(ctx, http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	newJWT, err := auth.RefreshToken(request.RefreshToken)
	if err != nil {
		sendErr(ctx, http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	sendSuccess(ctx, gin.H{
		"jwt_token": newJWT,
	})
}

// @Summary      User Logout
// @Description  Logs out the user and invalidates token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /auth/logout [post]
func LogoutHandler(ctx *gin.Context) {
	sendSuccess(ctx, "Logout!")
}
