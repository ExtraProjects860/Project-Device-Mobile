package repository

import (
	"context"

	"gorm.io/gorm"
)

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

func getByPagination[T any](db *gorm.DB, itemsPerPage, currentPage uint) ([]T, uint, uint, error) {
	var models []T

	offset, totalPages, totalItems := pagination(db, itemsPerPage, currentPage)

	err := db.
		Limit(int(itemsPerPage)).
		Offset(int(offset)).
		Find(&models).Error

	return models, totalPages, totalItems, err
}
