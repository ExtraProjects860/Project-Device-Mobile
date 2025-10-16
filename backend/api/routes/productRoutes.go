package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func registerProductRoutes(rg *gin.RouterGroup, appCtx *appcontext.AppContext) {
	{
		rg.GET("/products", handler.GetProductsHandler(
			appCtx,
			config.NewLogger("GET - PRODUCTS"),
		))

		rg.POST("/product", handler.CreateProductHandler(
			appCtx,
			config.NewLogger("POST - PRODUCT"),
		))

		rg.PATCH("/product", handler.UpdateProductHandler(
			appCtx,
			config.NewLogger("PATCH - PRODUCT"),
		))
	}
}
