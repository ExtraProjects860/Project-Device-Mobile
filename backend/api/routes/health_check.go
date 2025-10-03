package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitHealthCheckRoutes(r *gin.Engine, db *gorm.DB) {
	health := r.Group("/health")
	{
		health.GET("/api", handler.ApiHandler)

		health.GET("/database", func(ctx *gin.Context) {
			handler.DatabaseHandler(ctx, db)
		})

		health.GET("/emailservice", func(ctx *gin.Context) {
			serverDomain := config.GetEnv().API.EmailService
			handler.EmailServiceHandler(ctx, serverDomain)
		})
	}
}
