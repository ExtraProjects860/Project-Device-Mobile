package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/request-token", handler.RequestTokenHandler(
			appCtx,
			config.NewLogger("POST - AUTH-REQUEST"),
		))

		authGroup.POST("/reset-password", handler.ResetPasswordHandler(
			appCtx,
			config.NewLogger("POST - AUTH-RESET-PASSWORD"),
		))

		authGroup.POST("/reset-pass-log-in",
			middleware.JWTMiddleware(appCtx, config.NewLogger("MIDDLEWARE - JWT")),
			handler.ResetPasswordLogInHandler(
				appCtx,
				config.NewLogger("POST - AUTH-REST-PASS-LOG-IN"),
			))

		authGroup.POST("/login", handler.LoginHandler(
			appCtx,
			config.NewLogger("POST - AUTH-LOGIN"),
		))
	}
}
