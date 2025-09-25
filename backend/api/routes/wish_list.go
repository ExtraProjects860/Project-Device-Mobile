package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func RegisterWishListRoutes(rg *gin.RouterGroup) {
	{
		rg.GET("/wishlists", handler.GetItensWishListHandler)

		rg.POST("/wishlist", handler.AddInWishListHandler)

		rg.PATCH("/wishlist", handler.UpdateProductHandler)
	}
}
