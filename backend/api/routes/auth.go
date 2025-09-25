package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/request", handler.RequestTokenHandler)

		authGroup.POST("/reset", handler.ResetPasswordHandler)

		authGroup.POST("/login", handler.LoginHandler)

		authGroup.POST("/refresh-token", handler.RefreshTokenHandler)

		authGroup.POST("/logout", handler.LogoutHandler)
	}
}
