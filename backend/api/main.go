package main

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/router"
)

// @title Project Device Mobile API
// @version 1.0
// @description Essa é uma api voltada para um projeto extensionista para programação em dispositivos móveis
// @host localhost:5050
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}

	router.InitializeRouter(config.GetDB())
}
