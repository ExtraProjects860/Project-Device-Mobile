package repository

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
)

var logger *config.Logger

func init() {
	logger = config.GetLogger("repository")
}
