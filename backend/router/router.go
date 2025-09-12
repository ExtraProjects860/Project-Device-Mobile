package router

import "github.com/gin-gonic/gin"

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()

	initPingRoute(router)
	initMainRoutes(router)
	initSwaggerRoute(router)

	return
}
