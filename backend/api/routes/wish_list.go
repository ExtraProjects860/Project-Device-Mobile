package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"

	"github.com/gin-gonic/gin"
)

func registerWishListRoutes(rg *gin.RouterGroup, repo repository.WishListRepository) {
	wishlistHandler := handler.NewWishListHandler(repo)
	{
		rg.GET("/wishlist", wishlistHandler.GetWishListByUserIDHandler)

		rg.POST("/wishlist", wishlistHandler.AddInWishListHandler)

		rg.PATCH("/wishlist", wishlistHandler.UpdateWishListHandler)
	}
}
