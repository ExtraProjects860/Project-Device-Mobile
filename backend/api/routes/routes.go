package routes

import (
	docs "github.com/ExtraProjects860/Project-Device-Mobile/docs"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

const basePath string = "/api/v1"

func InitMainRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = basePath

	handler.InitializeHandler()

	api := r.Group(basePath)
	RegisterUserRoutes(api)
	RegisterProductRoutes(api)
	RegisterAuthRoutes(api)
	RegisterWishListRoutes(api)
}
