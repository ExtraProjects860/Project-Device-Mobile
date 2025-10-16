package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func registerRoleRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	{
		rg.GET("/roles", handler.GetRoles(
			appCtx,
			config.NewLogger("GET - ROLES"),
		))

		rg.POST("/role", handler.CreateRole(
			appCtx,
			config.NewLogger("POST - ROLE"),
		))

		rg.PUT("/role", handler.UpdateRole(
			appCtx,
			config.NewLogger("PUT - ROLE"),
		))
	}
}
