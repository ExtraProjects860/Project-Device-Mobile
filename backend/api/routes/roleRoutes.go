package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func registerRoleRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	logger := config.NewLogger("MIDDLEWARE")

	{
		rg.GET("/roles",
			middleware.JWTMiddleware(appCtx, logger),
			handler.GetRoles(
				appCtx,
				config.NewLogger("GET - ROLES"),
			))

		rg.POST("/role",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			handler.CreateRole(
				appCtx,
				config.NewLogger("POST - ROLE"),
			))

		rg.PUT("/role",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			handler.UpdateRole(
				appCtx,
				config.NewLogger("PUT - ROLE"),
			))
	}
}
