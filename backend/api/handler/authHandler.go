package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/auth"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
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
// @Param        email query string true "Email to search user"
// @Success      200 {object} string "Email to change Password Sent!"
// @Router       /api/v1/auth/request-token [post]
func RequestTokenHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.Query("email")
		if email == "" {
			logger.Info("email not send")
			response.SendErr(ctx, http.StatusBadRequest, errors.New("email not send, try again"))
			return
		}

		userService := service.GetUserService(appCtx)

		user, err := userService.GetByEmail(ctx, email)
		if err != nil {
			logger.Errorf("error to find user by email: %v", err)
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to find user by email. Try again or change email"))
			return
		}

		authService := service.GetAuthService(appCtx)

		err = authService.CreateToken(ctx, appCtx.Env.API.EmailService, user)
		if err != nil {
			logger.Errorf("error to generate token: %v", err)
			response.SendErr(ctx, http.StatusInternalServerError, err)
			return 
		}

		response.SendSuccess(ctx, http.StatusOK, "Email to change Password Sent!")
	}
}

// @Summary      Reset Password
// @Description  Resets user password using the token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        email query string true "Email to search user"
// @Param        token query string true "Token to change user password"
// @Param        request body request.ChangePassword true "Request body"
// @Success      200 {object} string "Change Password Successfully!"
// @Router       /api/v1/auth/reset-password [post]
func ResetPasswordHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.Query("email")
		if email == "" {
			logger.Info("email not send")
			response.SendErr(ctx, http.StatusBadRequest, errors.New("email not send, try again"))
			return
		}

		token := ctx.Query("token")
		if token == "" {
			logger.Info("token not send")
			response.SendErr(ctx, http.StatusBadRequest, errors.New("token not send, try again"))
			return
		}

		var input request.ChangePassword
		if err := request.ReadBodyJSON(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, errors.New("invalid input"))
			return
		}

		if input.NewPassword == "" {
			logger.Info("new_password can't be empty")
			response.SendErr(ctx, http.StatusBadRequest, errors.New("new_password can't be empty"))
			return
		}

		userService := service.GetUserService(appCtx)

		user, err := userService.GetByEmail(ctx, email)
		if err != nil {
			logger.Errorf("error to find user by email: %v", err)
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to find user by email. Try again or change email"))
			return
		}

		authService := service.GetAuthService(appCtx)
		err = authService.ChangePassword(
			ctx,
			appCtx.Env.API.EmailService,
			token,
			input.NewPassword,
			user,
			repository.NewPostgresUserRepository(appCtx.DB),
		)
		if err != nil {
			logger.Errorf("error to change user password: %v", err)
			response.SendErr(ctx, http.StatusInternalServerError, err)
			return
		}

		response.SendSuccess(ctx, http.StatusOK, "Change Password Successfully!")
	}
}

// @Summary      Reset Password Log In
// @Description  Resets user password log in system
// @Tags         auth
// @Param        request body request.ChangePassword true "Request body"
// @Accept       json
// @Produce      json
// @Success      200 {object} string "Change Password Successfully!"
// @Router       /api/v1/auth/reset-pass-log-in [post]
func ResetPasswordLogInHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := request.GetIdByToken(ctx)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnauthorized, err)
			return 
		}

		var input request.ChangePassword
		if err := request.ReadBodyJSON(ctx, &input); err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusUnprocessableEntity, errors.New("invalid input"))
			return
		}

		if input.NewPassword == "" {
			logger.Error()
			response.SendErr(ctx, http.StatusBadRequest, errors.New("new_password can't be empty"))
			return
		}

		authService := service.GetAuthService(appCtx)
		err = authService.ResetPassword(
			ctx,
			input.NewPassword,
			uint(uid),
			repository.NewPostgresUserRepository(appCtx.DB),
		)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to change password"))
			return
		}

		response.SendSuccess(ctx, http.StatusOK, "Change Password Successfully!")
	}
}

// @Summary      User Login
// @Description  Authenticates user and returns access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body request.LoginRequest true "Request body"
// @Success      201 {object} response.TokenResponse
// @Failure      400 {object} response.ErrResponse
// @Failure      401 {object} response.ErrResponse
// @Failure      422 {object} response.ErrResponse
// @Failure      500 {object} response.ErrResponse
// @Router       /api/v1/auth/login [post]
func LoginHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input request.LoginRequest
		if err := request.ReadBodyJSON(ctx, &input); err != nil {
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

		// TODO o token foi simplificado, porém um dia pode ter refresh
		accessToken, err := auth.GenerateAccessToken(
			userID,
			appCtx.Env.API.JwtKey,
			*input.RememberMe,
		)
		if err != nil {
			logger.Error(err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("error to generate jwt token and refresh token"))
			return
		}

		response.SendSuccess(ctx, http.StatusCreated, response.TokenResponse{
			Access: accessToken,
		})
	}
}
