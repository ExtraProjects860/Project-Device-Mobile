package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"

	"github.com/gin-gonic/gin"
)

func RegisterWishListRoutes(rg *gin.RouterGroup) {
	repo := repository.NewPostgresWishListRepository()
	{
		rg.GET("/wishlist", func(ctx *gin.Context) {
			handler.GetWishListByUserIDHandler(ctx, repo)
		})

		rg.POST("/wishlist", func(ctx *gin.Context) {
			handler.AddInWishListHandler(ctx, repo)
		})

		rg.PATCH("/wishlist", func(ctx *gin.Context) {
			handler.UpdateWishListHandler(ctx, repo)
		})
	}
}
