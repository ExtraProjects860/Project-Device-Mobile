package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/auth"
	"github.com/gin-gonic/gin"
)

// @Summary      Request Password Token
// @Description  Requests a reset token for user password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /api/v1/auth/request-token [post]
func RequestTokenHandler(ctx *gin.Context) {
	sendSuccess(ctx, http.StatusOK, "Require Password!")
}

// @Summary      Reset Password
// @Description  Resets user password using the token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /api/v1/auth/reset-password [post]
func ResetPasswordHandler(ctx *gin.Context) {
	sendSuccess(ctx, http.StatusOK, "Change Password!")
}

// @Summary      User Login
// @Description  Authenticates user and returns access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body LoginRequest true "Request body"
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /api/v1/auth/login [post]
func LoginHandler(ctx *gin.Context) {
	var request LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	// TODO Simulating user authentication (replace with real logic) *se não vai dar merda, e precisa implementar o repository*
	if request.Email != "test@gmail.com" || request.Password != "1234ok" {
		sendErr(ctx, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	user := &UserData{ID: 1, Email: "test@gmail.com"}
	jwtToken, refreshToken, err := auth.GenerateTokens(user.ID)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, errors.New("error to generate jwt token and refresh token"))
		return
	}

	sendSuccess(ctx, http.StatusCreated, gin.H{
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
// @Router       /api/v1/auth/refresh-token [post]
func RefreshTokenHandler(ctx *gin.Context) {
	var request auth.RequestRefresh
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	newJWT, err := auth.RefreshToken(request.RefreshToken)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, errors.New("error to generate new jwt token"))
		return
	}

	sendSuccess(ctx, http.StatusOK, gin.H{
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
// @Router       /api/v1/auth/logout [post]
func LogoutHandler(ctx *gin.Context) {
	sendSuccess(ctx, http.StatusOK, "Logout!")
}
