package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	logger := config.NewLogger("MIDDLEWARE")

	{
		rg.GET("/users",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			handler.GetUsersHandler(
				appCtx, config.NewLogger("GET - USERS"),
			))

		rg.GET("/user",
			middleware.JWTMiddleware(appCtx, logger),
			handler.GetInfoUserHandler(
				appCtx, config.NewLogger("GET - USER"),
			))

		rg.POST("/user",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			middleware.ImageOptional(appCtx, logger),
			handler.CreateUserHandler(
				appCtx, config.NewLogger("POST - USERS"),
			))

		/*
			TODO verificar depois de fazer outra rota em auth, porém para o usuário atualizar os dados
		*/
		rg.PATCH("/user",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			middleware.ImageOptional(appCtx, logger),
			handler.UpdateUserHandler(
				appCtx, config.NewLogger("PATCH - USERS"),
			))
	}
}
