package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func RegisterPromotionRoutes(rg *gin.RouterGroup) {
	promotionGroup := rg.Group("/promotions")
	promotionGroup.GET("/", handler.GetPromotionsHandler)
	promotionGroup.POST("/", handler.CreatePromotionHandler)
	promotionGroup.PATCH("/{id}", handler.UpdatePromotionHandler)
}
