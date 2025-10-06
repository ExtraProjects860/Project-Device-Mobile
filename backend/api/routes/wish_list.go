package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func registerWishListRoutes(rg *gin.RouterGroup) {
	{
		rg.GET("/wishlist", handler.GetWishListByUserIDHandler)

		rg.POST("/wishlist", handler.AddInWishListHandler)

		rg.PATCH("/wishlist", handler.UpdateWishListHandler)
	}
}
