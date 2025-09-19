package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func InitHealthCheckRoutes(r *gin.Engine) {
	const basePath = "/health"
	health := r.Group(basePath)
	{
		health.GET("/api", handler.Api)

		health.GET("/database", handler.Database)

		health.GET("/emailservice", handler.EmailService)
	}
}
