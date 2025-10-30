package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func registerProductRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	logger := config.NewLogger("MIDDLEWARE")
	{
		rg.GET("/products",
			middleware.JWTMiddleware(appCtx, logger),
			handler.GetProductsHandler(
				appCtx,
				config.NewLogger("GET - PRODUCTS"),
			))

		rg.POST("/product",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			middleware.ImageOptional(appCtx, logger),
			handler.CreateProductHandler(
				appCtx,
				config.NewLogger("POST - PRODUCT"),
			))

		rg.PATCH("/product",
			middleware.JWTMiddleware(appCtx, logger),
			middleware.AdminPermission(appCtx, logger),
			middleware.ImageOptional(appCtx, logger),
			handler.UpdateProductHandler(
				appCtx,
				config.NewLogger("PATCH - PRODUCT"),
			))
	}
}
