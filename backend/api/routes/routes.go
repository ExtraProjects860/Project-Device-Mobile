package routes

import (
<<<<<<< HEAD
	docs "github.com/ExtraProjects860/Project-Device-Mobile/docs"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
=======
>>>>>>> dev
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

<<<<<<< HEAD
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
=======
func InitRoutesApiV1(router *gin.Engine, db *gorm.DB) {
	apiV1 := router.Group("/api/v1")
	
	RegisterUserRoutes(apiV1, repository.NewPostgresUserRepository(db))
	RegisterProductRoutes(apiV1, repository.NewPostgresProductRepository(db))
	RegisterAuthRoutes(apiV1)
	RegisterWishListRoutes(apiV1, repository.NewPostgresWishListRepository(db))
>>>>>>> dev
}
