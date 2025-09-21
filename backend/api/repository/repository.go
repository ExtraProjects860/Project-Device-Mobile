package repository

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeRepository() {
	db = config.GetDB()
	logger = config.GetLogger("repository")
}
