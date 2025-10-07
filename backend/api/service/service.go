package service

import "github.com/ExtraProjects860/Project-Device-Mobile/config"

// TODO aplicar DTOs em service e separar a lógica do handler

var logger *config.Logger

func init() {
	logger = config.GetLogger("service")
}
