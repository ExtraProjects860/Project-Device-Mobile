package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	authGroup.POST("/request", handler.RequestToken)
	authGroup.POST("/reset", handler.ResetPassword)
	authGroup.POST("/login", handler.Login)
	authGroup.POST("/refresh-token", handler.RefreshToken)
	authGroup.POST("/logout", handler.Logout)
}
