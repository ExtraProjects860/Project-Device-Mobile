package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutesApiV1(router *gin.Engine, db *gorm.DB) {
	apiV1 := router.Group("/api/v1")
	
	registerUserRoutes(apiV1, repository.NewPostgresUserRepository(db))
	registerProductRoutes(apiV1, repository.NewPostgresProductRepository(db))
	registerAuthRoutes(apiV1)
	registerWishListRoutes(apiV1, repository.NewPostgresWishListRepository(db))
}
