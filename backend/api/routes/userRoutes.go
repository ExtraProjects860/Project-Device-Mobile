package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

// TODO colocar middlewares dps
func registerUserRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	{
		rg.GET("/users", handler.GetUsersHandler(
			appCtx, config.NewLogger("GET - USERS"),
		))

		rg.GET("/user", handler.GetInfoUserHandler(
			appCtx, config.NewLogger("GET - USER"),
		))

		rg.POST("/user", handler.CreateUserHandler(
			appCtx, config.NewLogger("POST - USERS"),
		))

		rg.PATCH("/user", handler.UpdateUserHandler(
			appCtx, config.NewLogger("PATCH - USERS"),
		))
	}
}
