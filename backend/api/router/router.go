package router

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/routes"
	"github.com/gin-gonic/gin"
)

func configureNetwork(router *gin.Engine) {
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{})
}

func InitializeRouter() {
	router := gin.Default()

	configureNetwork(router)

	routes.InitHealthCheckRoutes(router)
	routes.InitRoutesApiV1(router)
	routes.InitSwaggerRoute(router)

	port := config.GetEnv().API.Port
	router.Run(fmt.Sprintf(":%v", port))
}
