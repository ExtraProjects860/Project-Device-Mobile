package repository

import (
	"math"

	"gorm.io/gorm"
)

type PaginationDTO struct {
	Data        any  `json:"data"`
	CurrentPage uint `json:"current_page"`
	TotalPages  uint `json:"total_pages"`
<<<<<<< HEAD
=======
	TotalItems  uint `json:"total_items"`
>>>>>>> dev
}

func count(query *gorm.DB) uint {
	var count int64
	query.Count(&count)

	return uint(count)
}

func validationPagination(currentPage *uint, itemsPerPage *uint) {
	if *currentPage == 0 {
		*currentPage = 1
	}

	if *itemsPerPage == 0 {
		*itemsPerPage = 1
	}
}

<<<<<<< HEAD
func pagination(query *gorm.DB, itemsPerPage uint, currentPage uint) (uint, uint) {
	validationPagination(&currentPage, &itemsPerPage)
	lengthItems := count(query)
	if lengthItems == 0 {
		return 0, 1
=======
func pagination(query *gorm.DB, itemsPerPage uint, currentPage uint) (uint, uint, uint) {
	validationPagination(&currentPage, &itemsPerPage)
	lengthItems := count(query)
	if lengthItems == 0 {
		return 0, 1, 0
>>>>>>> dev
	}

	totalPages := uint(math.Ceil(float64(lengthItems) / float64(itemsPerPage)))

	if currentPage > totalPages {
		currentPage = totalPages
	}

	paginationOffset := (currentPage - 1) * itemsPerPage

<<<<<<< HEAD
	return paginationOffset, totalPages
=======
	return paginationOffset, totalPages, lengthItems
>>>>>>> dev
}
