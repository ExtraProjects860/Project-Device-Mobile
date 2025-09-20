package handler

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// @Summary      Request Password Token
// @Description  Requests a reset token for user password
// @Tags         auth
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /auth/request-token [post]
func RequestToken(ctx *gin.Context) {
	sendSuccess(ctx, "Require Password!")
}

// @Summary      Reset Password
// @Description  Resets user password using the token
// @Tags         auth
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /auth/reset-password [post]
func ResetPassword(ctx *gin.Context) {
	sendSuccess(ctx, "Change Password!")
}

// @Summary      User Login
// @Description  Authenticates user and returns access token
// @Tags         auth
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /auth/login [post]
func Login(ctx *gin.Context) {
	sendSuccess(ctx, "Login!")
}

// @Summary      Refresh Token
// @Description  Refreshes the authentication token
// @Tags         auth
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /auth/refresh-token [post]
func RefreshToken(ctx *gin.Context) {
	sendSuccess(ctx, "Refresh Token!")
}

// @Summary      User Logout
// @Description  Logs out the user and invalidates token
// @Tags         auth
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /auth/logout [post]
func Logout(ctx *gin.Context) {
	sendSuccess(ctx, "Logout!")
}
