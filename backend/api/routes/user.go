package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")
	userGroup.GET("/", handler.GetUsers)
	userGroup.GET("/{id}", handler.GetInfoUser)
	userGroup.POST("/", handler.CreateUser)
	userGroup.PATCH("/{id}", handler.UpdateUser)
}
