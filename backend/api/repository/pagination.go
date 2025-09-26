package repository

import "math"

type PaginationDTO struct {
	Data        any
	CurrentPage uint
	TotalPages  uint
}

func count(model any) uint {
	var count int64
	db.Model(model).Count(&count)

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

func pagination(model any, itemsPerPage uint, currentPage uint) (uint, uint) {
	validationPagination(&currentPage, &itemsPerPage)
	lengthItems := count(model)
	if lengthItems == 0 {
		return 0, 1
	}

	totalPages := uint(math.Ceil(float64(lengthItems) / float64(itemsPerPage)))

	if currentPage > totalPages {
		currentPage = totalPages
	}

	paginationOffset := (currentPage - 1) * itemsPerPage

	return paginationOffset, totalPages
}
