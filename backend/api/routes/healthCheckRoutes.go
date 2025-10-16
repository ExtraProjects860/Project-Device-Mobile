package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler"
	"github.com/gin-gonic/gin"
)

func InitHealthCheckRoutes(r *gin.Engine, appCtx *appcontext.AppContext) {
	health := r.Group("/health")
	logger := config.NewLogger("GET - HEALTHCHECK")

	{
		health.GET("/api", handler.ApiHandler)

		health.GET("/database", handler.DatabaseHandler(
			appCtx,
			logger,
		))

		health.GET("/emailservice", handler.EmailServiceHandler(
			appCtx,
			logger,
		))
	}
}
