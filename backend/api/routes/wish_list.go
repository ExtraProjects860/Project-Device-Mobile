package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func RegisterWishListRoutes(rg *gin.RouterGroup) {
	wishListGroup := rg.Group("/wish-list")
	wishListGroup.GET("/", handler.GetItensWishListHandler)
	wishListGroup.POST("/", handler.AddProductWishListHandler)
	
	// TODO revisar um método melhor pra essa rota
	wishListGroup.PATCH("/{id}", handler.DeleteProductWishListHandler)
}
