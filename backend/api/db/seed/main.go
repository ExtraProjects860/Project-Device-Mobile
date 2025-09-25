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

func initializeSeed() {
	if err := config.Init(); err != nil {
		panic(fmt.Errorf("failed to init config: %v", err))
	}
	logger = config.GetLogger("migrate")
	db = config.GetDB()
}

func main() {
	initializeSeed()
	resetDB()
	seeds()
}
