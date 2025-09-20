package handler

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
	env *config.EnvVariables
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
	env = config.GetEnv()
}
