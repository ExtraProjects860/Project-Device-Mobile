package seed

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler() {
	if err := config.Init(); err != nil {
		panic(fmt.Errorf("failed to init config: %v", err))
	}
	logger = config.GetLogger("seed")
	db = config.GetDB()
}
