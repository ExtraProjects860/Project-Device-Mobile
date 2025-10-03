package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
func RegisterProductRoutes(rg *gin.RouterGroup) {
	repo := repository.NewPostgresProductRepository()
	{
		rg.GET("/products", func(ctx *gin.Context) {
			handler.GetProductsHandler(ctx, repo)
		})

		rg.POST("/product", func(ctx *gin.Context) {
			handler.CreateProductHandler(ctx, repo)
		})

		rg.PATCH("/product", func(ctx *gin.Context) {
			handler.UpdateProductHandler(ctx, repo)
		})
=======
func RegisterProductRoutes(rg *gin.RouterGroup, repo repository.ProductRepository) {
	productHandler := handler.NewProductHandler(repo)
	{
		rg.GET("/products", productHandler.GetProductsHandler)

		rg.POST("/product", productHandler.CreateProductHandler)

		rg.PATCH("/product", productHandler.UpdateProductHandler)
>>>>>>> dev
	}
}
