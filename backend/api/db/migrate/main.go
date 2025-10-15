package main

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
)

func main() {
	logger := config.NewLogger("migrate")

	_, db, err := config.Init()
	if err != nil {
		panic(fmt.Errorf("failed to init config: %v", err))
	}

	if err := migrateDatabase(db, logger); err != nil {
		logger.Errorf("AutoMigrate error: %v", err)
		panic(fmt.Errorf("failed migrate models to database: %v", err))
	}
}
