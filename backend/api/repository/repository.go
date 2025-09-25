package repository

import (
	"math"

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

func pagination(legthItens uint, itemsPerPage uint, currentPage uint) map[string]uint {
	currentPage = uint(math.Max(0, float64(currentPage)))
	
	var totalPages uint
	if itemsPerPage > 0 {
		totalPages = uint(math.Floor(float64(legthItens) / float64(itemsPerPage)))
	} else {
		totalPages = 0
	}

	return map[string]uint{
		"currentPage": currentPage,
		"totalPages": totalPages,
	}
}
