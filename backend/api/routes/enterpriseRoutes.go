package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func registerEnterpriseRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	logger := config.NewLogger("MIDDLEWARE")

	{
		rg.GET("/enterprises",
			middleware.JWTMiddleware(appCtx, logger),
			handler.GetEnterprises(
				appCtx,
				config.NewLogger("GET - ENTERPRISES"),
			))

		rg.POST("/enterprise",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			handler.CreateEnterprise(
				appCtx,
				config.NewLogger("POST - ENTERPRISES"),
			))

		rg.PUT("/enterprise",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			handler.UpdateEnterprise(
				appCtx,
				config.NewLogger("PUT - ENTERPRISES"),
			))
	}
}
