package main

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
)

var logger *config.Logger = config.GetLogger("migrate")

func main() {
	if err := config.Init(); err != nil {
		panic(fmt.Errorf("failed to init config: %v", err))
	}

	if err := migrateDatabase(config.GetDB()); err != nil {
		logger.Errorf("AutoMigrate error: %v", err)
		panic(fmt.Errorf("failed migrate models to database: %v", err))
	}
}
