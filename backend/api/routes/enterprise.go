package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func registerEnterpriseRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	{
		rg.GET("/enterprises", handler.GetEnterprises(
			appCtx,
			config.NewLogger("GET - ENTERPRISES"),
		))

		rg.POST("/enterprise", handler.CreateEnterprise(
			appCtx,
			config.NewLogger("POST - ENTERPRISES"),
		))

		rg.PUT("/enterprise", handler.UpdateEnterprise(
			appCtx,
			config.NewLogger("PUT - ENTERPRISES"),
		))
	}
}
