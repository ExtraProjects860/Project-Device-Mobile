package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func registerWishListRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	logger := config.NewLogger("MIDDLEWARE")

	{
		rg.GET("/wishlist",
			middleware.JWTMiddleware(appCtx, logger),
			handler.GetWishListByUserIDHandler(
				appCtx,
				config.NewLogger("GET - WISHLIST"),
			))

		rg.POST("/wishlist",
			middleware.JWTMiddleware(appCtx, logger),
			handler.AddInWishListHandler(
				appCtx,
				config.NewLogger("POST - WISHLIST"),
			))

		rg.DELETE("/wishlist",
			middleware.JWTMiddleware(appCtx, logger),
			handler.DeleteInWishListHandler(
				appCtx,
				config.NewLogger("DELETE - WISHLIST"),
			))
	}
}
