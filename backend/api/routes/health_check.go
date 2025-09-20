package routes

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func InitHealthCheckRoutes(r *gin.Engine) {
	health := r.Group(fmt.Sprintf("%v/health", basePath))
	{
		health.GET("/api", handler.Api)

		health.GET("/database", handler.Database)

		health.GET("/emailservice", handler.EmailService)
	}
}
