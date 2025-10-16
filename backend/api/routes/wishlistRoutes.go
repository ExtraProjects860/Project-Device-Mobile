package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func registerWishListRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	{
		rg.GET("/wishlist", handler.GetWishListByUserIDHandler(
			appCtx,
			config.NewLogger("GET - WISHLIST"),
		))

		rg.POST("/wishlist", handler.AddInWishListHandler(
			appCtx,
			config.NewLogger("POST - WISHLIST"),
		))

		rg.DELETE("/wishlist", handler.DeleteInWishListHandler(
			appCtx,
			config.NewLogger("DELETE - WISHLIST"),
		))
	}
}
