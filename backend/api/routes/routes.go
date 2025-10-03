package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutesApiV1(router *gin.Engine, db *gorm.DB) {
	apiV1 := router.Group("/api/v1")
	
	RegisterUserRoutes(apiV1, repository.NewPostgresUserRepository(db))
	RegisterProductRoutes(apiV1, repository.NewPostgresProductRepository(db))
	RegisterAuthRoutes(apiV1)
	RegisterWishListRoutes(apiV1, repository.NewPostgresWishListRepository(db))
}
