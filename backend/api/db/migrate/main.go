package main

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func initializeMigrate() {
	if err := config.Init(); err != nil {
		panic(fmt.Errorf("failed to init config: %v", err))
	}
	logger = config.GetLogger("migrate")
	db = config.GetDB()
}

func main() {
	initializeMigrate()

	if err := migrateDatabase(db); err != nil {
		logger.Errorf("AutoMigrate error: %v", err)
		panic(fmt.Errorf("failed migrate models to database: %v", err))
	}
}
