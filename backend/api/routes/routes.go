package routes

import (
	docs "github.com/ExtraProjects860/Project-Device-Mobile/docs"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

const basePath string = "/api/v1"

var (

)

func InitMainRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = basePath

	repository.InitializeRepository()
	handler.InitializeHandler()

	api := r.Group(basePath)
	RegisterUserRoutes(api)
	RegisterProductRoutes(api)
	RegisterAuthRoutes(api)
	RegisterWishListRoutes(api)
}
