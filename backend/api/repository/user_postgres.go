package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user schemas.User) error
	GetInfoUser(id uint) (schemas.User, error)
	GetUsers() ([]schemas.User, error)
	UpdateUser(id uint) (schemas.User, error)
}

// TODO colocar mensagens nos logs aqui depois

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user schemas.User) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			logger.Errorf("%v", err)
			return err
		}
		return nil
	})
}

func (r *PostgresUserRepository) GetInfoUser(ctx context.Context, id uint) (schemas.User, error) {
	var user schemas.User
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, id).Error; err != nil {
			logger.Errorf("%v", err)
			return err
		}

		return nil
	})

	return user, err
}

func (r *PostgresUserRepository) GetUsers(ctx context.Context) ([]schemas.User, error) {
	var user []schemas.User
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Preload("Role").
			Preload("Enterprise").
			Find(&user).Error; err != nil {
			logger.Errorf("%v", err)
			return err
		}

		return nil
	})

	return user, err
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, id uint, user schemas.User) (schemas.User, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&schemas.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
			logger.Errorf("%v", err)
			return err
		}

		return nil
	})

	return user, err
}
