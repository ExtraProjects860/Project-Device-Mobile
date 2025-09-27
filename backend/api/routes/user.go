package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

// TODO colocar middlewares dps
func RegisterUserRoutes(rg *gin.RouterGroup) {
	repo := repository.NewPostgresUserRepository()
	{
		rg.GET("/users", func(ctx *gin.Context) {
			handler.GetUsersHandler(ctx, repo)
		})

		rg.GET("/user", func(ctx *gin.Context) {
			handler.GetInfoUserHandler(ctx, repo)
		})

		rg.POST("/user", func(ctx *gin.Context) {
			handler.CreateUserHandler(ctx, repo)
		})

		rg.PATCH("/user", func(ctx *gin.Context) {
			handler.UpdateUserHandler(ctx, repo)
		})
	}
}
