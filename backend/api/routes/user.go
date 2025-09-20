package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")
	userGroup.GET("/", middleware.JWTMiddleware(), handler.GetUsersHandler)
	userGroup.GET("/{id}", handler.GetInfoUserHandler)
	userGroup.POST("/", handler.CreateUserHandler)
	userGroup.PATCH("/{id}", handler.UpdateUserHandler)
}
