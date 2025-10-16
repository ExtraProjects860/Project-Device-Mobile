package main

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/routes"
	"gorm.io/gorm"
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

func setupContext(env *config.EnvVariables, db *gorm.DB) *appcontext.AppContext {
	return &appcontext.AppContext{
		Env: env,
		DB:  db,
	}
}

func main() {
	logger := config.NewLogger("main")

	env, db, err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}

	appCtx := setupContext(env, db)

	r := routes.InitializeRouter(appCtx)
	r.Run(fmt.Sprintf(":%v", env.API.Port))
}
