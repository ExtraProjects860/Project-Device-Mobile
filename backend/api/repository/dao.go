package repository

import (
	"context"
	"reflect"
	"strings"
	"sync"

	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"gorm.io/gorm"
)

var searchableFieldsCache sync.Map

func firstWhere[T any](db *gorm.DB, query string, args ...any) (T, error) {
	var model T
	err := db.Where(query, args...).First(&model).Error
	return model, err
}

func getByID[T any](db *gorm.DB, id uint) (T, error) {
	var model T
	err := db.First(&model, id).Error
	return model, err
}

func create[T any](ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(new(T)).Create(entity).Error
	})
}

func updateByID[T any](ctx context.Context, db *gorm.DB, entity *T, id uint) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(new(T)).Where("id = ?", id).Updates(entity).Error
	})
}

func update[T any](ctx context.Context, db *gorm.DB, entity *T, queryWhere string, args ...any) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(new(T)).Where(queryWhere, args...).Updates(entity).Error
	})
}

func deleteByID[T any](ctx context.Context, db *gorm.DB, entity *T, id uint) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(new(T)).Where("id = ?", id).Delete(entity).Error
	})
}

func delete[T any](ctx context.Context, db *gorm.DB, entity *T, queryWhere string, args ...any) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(new(T)).Where(queryWhere, args...).Delete(entity).Error
	})
}

func filterSearch[T any](query *gorm.DB, search string) *gorm.DB {
	if search == "" {
		return query
	}

	var fields []string
	typ := reflect.TypeOf((*T)(nil)).Elem()

	if cached, ok := searchableFieldsCache.Load(typ); ok {
		fields = cached.([]string)
	}

	var model T
	stmt := &gorm.Statement{DB: query}
	if err := stmt.Parse(&model); err != nil {
		return query
	}

	for _, field := range stmt.Schema.Fields {
		if field.DBName == "" ||
			field.DBName == "created_at" ||
			field.DBName == "updated_at" ||
			field.DBName == "deleted_at" {
			continue
		}

		if field.FieldType.Kind() != reflect.String {
			continue
		}

		fields = append(fields, field.DBName)

		searchableFieldsCache.Store(typ, fields)
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

func getByPagination[T any](db *gorm.DB, paginationSearch request.PaginationSearch) ([]T, uint, uint, error) {
	var models []T

	offset, totalPages, totalItems := pagination(
		db,
		paginationSearch.ItemsPerPage,
		paginationSearch.CurrentPage,
	)

	db = filterSearch[T](db, paginationSearch.SearchFilter)

	err := db.
		Order("id " + paginationSearch.ItemsOrder.String()).
		Limit(int(paginationSearch.ItemsPerPage)).
		Offset(int(offset)).
		Find(&models).Error

	return models, totalPages, totalItems, err
}
