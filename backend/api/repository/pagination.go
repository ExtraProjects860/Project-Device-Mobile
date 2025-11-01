package repository

import (
	"math"
	"reflect"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"gorm.io/gorm"
)

type SearchableModel interface {
	GetSearchableFields() []string
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

func pagination(query *gorm.DB, itemsPerPage uint, currentPage uint) (uint, uint, uint) {
	validationPagination(&currentPage, &itemsPerPage)
	lengthItems := count(query)
	if lengthItems == 0 {
		return 0, 1, 0
	}

	totalPages := uint(math.Ceil(float64(lengthItems) / float64(itemsPerPage)))

	if currentPage > totalPages {
		currentPage = totalPages
	}

	paginationOffset := (currentPage - 1) * itemsPerPage

	return paginationOffset, totalPages, lengthItems
}

func filterSearch[T SearchableModel](query *gorm.DB, search string) *gorm.DB {
	if search == "" {
		return query
	}

	var fields []string
	typ := reflect.TypeOf((*T)(nil)).Elem()

	if cached, ok := searchableFieldsCache.Load(typ); ok {
		fields = cached.([]string)
	} else {
		var model T
		fields = model.GetSearchableFields()
		searchableFieldsCache.Store(typ, fields)
	}

	if len(fields) == 0 {
		return query
	}

	var args []any
	var whereParts []string
	for _, col := range fields {
		whereParts = append(whereParts, col+" ILIKE ?")
		args = append(args, "%"+search+"%")
	}

	return query.Where(
		strings.Join(whereParts, " OR "),
		args...,
	)
}

func getByPagination[T SearchableModel](db *gorm.DB, paginationSearch request.PaginationSearch) ([]T, uint, uint, error) {
	var models []T

	db = filterSearch[T](db, paginationSearch.SearchFilter)

	offset, totalPages, totalItems := pagination(
		db,
		paginationSearch.ItemsPerPage,
		paginationSearch.CurrentPage,
	)

	err := db.
		Order("id " + paginationSearch.ItemsOrder.String()).
		Limit(int(paginationSearch.ItemsPerPage)).
		Offset(int(offset)).
		Find(&models).Error

	return models, totalPages, totalItems, err
}
