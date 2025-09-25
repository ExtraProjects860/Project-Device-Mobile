package routes

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func InitHealthCheckRoutes(r *gin.Engine) {
	health := r.Group(fmt.Sprintf("%v/health", basePath))
	{
		health.GET("/api", handler.ApiHandler)

		health.GET("/database", handler.DatabaseHandler)

		health.GET("/emailservice", handler.EmailServiceHandler)
	}
}
