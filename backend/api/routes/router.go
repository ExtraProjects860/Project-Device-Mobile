package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/middleware"
	"github.com/gin-gonic/gin"
)

func configureNetwork(router *gin.Engine) {
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{})
}

func InitializeRouter(appCtx *appcontext.AppContext) *gin.Engine {
	router := gin.Default()

	configureNetwork(router)

	middleware.SecurityHeaders(router)
	middleware.SetCors(router)

	InitHealthCheckRoutes(router, appCtx)
	InitRoutesApiV1(router, appCtx)
	InitSwaggerRoute(router)

	return router
}
