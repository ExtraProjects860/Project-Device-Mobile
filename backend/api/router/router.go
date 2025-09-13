package router

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/routes"
	"github.com/gin-gonic/gin"
)

func configureRouter(router *gin.Engine) {
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{})
}

func InitializeRouter()  {
	router := gin.Default()

	configureRouter(router)

	routes.InitHealthCheckRoutes(router)
	routes.InitMainRoutes(router)
	routes.InitSwaggerRoute(router)

	port := config.GetEnv().API.Port
	router.Run(fmt.Sprintf(":%v", port))
}
