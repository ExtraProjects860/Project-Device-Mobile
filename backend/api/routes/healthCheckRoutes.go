package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func InitHealthCheckRoutes(r *gin.Engine) {
	health := r.Group("/health")
	{
		health.GET("/api", handler.ApiHandler)

		health.GET("/database", func(ctx *gin.Context) {
			handler.DatabaseHandler(ctx, config.GetDB())
		})

		health.GET("/emailservice", func(ctx *gin.Context) {
			serverDomain := config.GetEnv().API.EmailService
			handler.EmailServiceHandler(ctx, serverDomain)
		})
	}
}
