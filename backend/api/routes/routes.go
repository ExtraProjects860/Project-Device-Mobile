package routes

import (
	"github.com/gin-gonic/gin"
	docs "github.com/ExtraProjects860/Project-Device-Mobile/docs"
)

const basePath string = "/api/v1"

func InitMainRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = basePath

	api := r.Group(basePath)
	RegisterUserRoutes(api)
	RegisterProductRoutes(api)
	RegisterPromotionRoutes(api)
	RegisterAuthRoutes(api)
	RegisterWishListRoutes(api)
}
