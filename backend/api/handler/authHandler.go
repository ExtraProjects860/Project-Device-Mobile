package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/auth"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/ExtraProjects860/Project-Device-Mobile/service"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

// @Summary      Request Password Token
// @Description  Requests a reset token for user password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /api/v1/auth/request-token [post]
func RequestTokenHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response.SendSuccess(ctx, http.StatusOK, "Require Password!")
	}
}

// @Summary      Reset Password
// @Description  Resets user password using the token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /api/v1/auth/reset-password [post]
func ResetPasswordHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response.SendSuccess(ctx, http.StatusOK, "Change Password!")
	}
}

// @Summary      Reset Password Log In
// @Description  Resets user password log in system
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /api/v1/auth/reset-pass-log-in [post]
func ResetPasswordLogInHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response.SendSuccess(ctx, http.StatusOK, "Change Password!")
	}
}

// @Summary      User Login
// @Description  Authenticates user and returns access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body request.LoginRequest true "Request body"
// @Success      201 {object} response.TokensResponse
// @Failure      400 {object} response.ErrResponse
// @Failure      401 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/auth/login [post]
func LoginHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var input request.LoginRequest
		if err := request.ReadBody(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, errors.New("invalid input"))
			return
		}

		if err := request.ValidateBodyReq(&input, ctx, utils.GetValidate()); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusBadRequest, err)
			return
		}

		authService := service.GetAuthService(appCtx)
		userID, err := authService.VerifyCredentials(
			ctx,
			input,
			repository.NewPostgresUserRepository(appCtx.DB),
		)

		if err != nil {
			response.SendErr(ctx, http.StatusUnauthorized, err)
			return
		}

		accessToken, refreshToken, err := auth.GenerateTokens(
			userID,
			appCtx.Env.API.JwtKey,
			appCtx.Env.API.RefreshKey,
		)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to generate jwt token and refresh token"))
			return
		}

		response.SendSuccess(ctx, http.StatusCreated, response.TokenResponse{
			Access:  accessToken,
			Refresh: refreshToken,
		})
	}
}

// @Summary      Refresh Token
// @Description  Refreshes the authentication token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      201 {object} response.TokenResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/auth/refresh-token [post]
func RefreshTokenHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input auth.RequestRefresh
		if err := request.ReadBody(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, errors.New("invalid input"))
			return
		}

		newAccessToken, err := auth.RefreshToken(
			input.RefreshToken,
			appCtx.Env.API.JwtKey,
			appCtx.Env.API.RefreshKey,
		)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to generate new jwt token"))
			return
		}

		response.SendSuccess(ctx, http.StatusCreated, response.TokenResponse{
			Access: newAccessToken,
		})
	}
}
