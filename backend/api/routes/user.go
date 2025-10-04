package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

// TODO colocar middlewares dps
func registerUserRoutes(rg *gin.RouterGroup, repo repository.UserRepository) {
	userHandler := handler.NewUserHandler(repo)
	{
		rg.GET("/users", userHandler.GetUsersHandler)

		rg.GET("/user", userHandler.GetInfoUserHandler)

		rg.POST("/user", userHandler.CreateUserHandler)

		rg.PATCH("/user", userHandler.UpdateUserHandler)
	}
}
