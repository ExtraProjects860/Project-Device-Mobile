package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup, repo repository.ProductRepository) {
	productHandler := handler.NewProductHandler(repo)
	{
		rg.GET("/products", productHandler.GetProductsHandler)

		rg.POST("/product", productHandler.CreateProductHandler)

		rg.PATCH("/product", productHandler.UpdateProductHandler)
	}
}
