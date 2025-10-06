package repository

import (
	"context"

	"gorm.io/gorm"
)

// TODO Adicionar delete na wishlist

func getByID[T any](ctx context.Context, query *gorm.DB, id uint) (T, error) {
	var model T
	err := query.First(&model, id).Error
	return model, err
}

func create[T any](ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(entity).Error
	})
}

// TODO alterar update para não ser genérico permitindo parcial update

func update[T any](ctx context.Context, db *gorm.DB, id uint, entity *T) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(entity).Where("id = ?", id).Updates(entity).Error
	})
}

func deleteByID[T any](ctx context.Context, db *gorm.DB, id uint) error {
	var entity T
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&entity).Error
	})
}

func getByPagination[T any](query *gorm.DB, itemsPerPage, currentPage uint) ([]T, uint, uint, error) {
	var models []T

	offset, totalPages, totalItems := pagination(query, itemsPerPage, currentPage)

	err := query.
		Limit(int(itemsPerPage)).
		Offset(int(offset)).
		Find(&models).Error

	return models, totalPages, totalItems, err
}
