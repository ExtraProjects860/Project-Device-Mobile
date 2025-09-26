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

type PaginationDTO struct {
	Data any;
	CurrentPage uint;
	TotalPages uint;
}

func count(model any) uint {
	var count int64
	db.Model(model).Count(&count)

	return uint(count)
}

func validationPagination(currentPage *uint, itemsPerPage *uint) {
	if *currentPage <= 0 {
		*currentPage = 1
	}

	if *itemsPerPage <= 0 {
		*itemsPerPage = 1
	}
}

func pagination(model any, itemsPerPage uint, currentPage uint) (uint, uint) {
	validationPagination(&currentPage, &itemsPerPage)
	legthItems := count(model)
	paginationOffset := (currentPage - 1) * itemsPerPage
	totalPages := uint(math.Ceil(float64(legthItems) / float64(itemsPerPage)))

	return paginationOffset, totalPages
}
