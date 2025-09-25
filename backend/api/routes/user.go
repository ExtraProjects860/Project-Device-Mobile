package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

// TODO colocar middlewares dps
func RegisterUserRoutes(rg *gin.RouterGroup) {
	{
		rg.GET("/users", handler.GetUsersHandler)

		rg.GET("/user", handler.GetInfoUserHandler)

		rg.POST("/user", handler.CreateUserHandler)

		rg.PATCH("/user", handler.UpdateUserHandler)
	}
}
