package handler

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
)

var (
	logger *config.Logger
	env *config.EnvVariables
)

func init() {
	logger = config.GetLogger("handler")
	env = config.GetEnv()
}
