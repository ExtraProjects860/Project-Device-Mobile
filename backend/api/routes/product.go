package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup) {
	productGroup := rg.Group("/products")
	productGroup.GET("/", handler.GetProducts)
	productGroup.POST("/", handler.CreateProduct)
	productGroup.PATCH("/{id}", handler.UpdateProduct)
}
