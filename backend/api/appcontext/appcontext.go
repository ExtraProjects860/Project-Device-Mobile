package appcontext

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"gorm.io/gorm"
)

type AppContext struct {
	Env *config.EnvVariables
	DB  *gorm.DB
}
