package router

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func configureNetwork(router *gin.Engine) {
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{})
}

func InitializeRouter(db *gorm.DB) {
	router := gin.Default()

	configureNetwork(router)

	routes.InitHealthCheckRoutes(router, db)
	routes.InitRoutesApiV1(router, db)
	routes.InitSwaggerRoute(router)

	port := config.GetEnv().API.Port
	router.Run(fmt.Sprintf(":%v", port))
}
