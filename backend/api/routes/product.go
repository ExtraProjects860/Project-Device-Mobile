package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func registerProductRoutes(rg *gin.RouterGroup) {
	{
		rg.GET("/products", handler.GetProductsHandler)

		rg.POST("/product", handler.CreateProductHandler)

		rg.PATCH("/product", handler.UpdateProductHandler)
	}
}
