package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/request", handler.RequestTokenHandler)

		authGroup.POST("/reset-password", handler.ResetPasswordHandler)

		authGroup.POST("/reset-pass-log-in", handler.ResetPasswordLogInHandler)

		authGroup.POST("/login", handler.LoginHandler)

		authGroup.POST("/refresh-token", handler.RefreshTokenHandler)
	}
}
