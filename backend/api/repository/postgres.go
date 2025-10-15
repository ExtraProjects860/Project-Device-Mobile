package repository

import (
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

type PostgresProductRepository struct {
	db *gorm.DB
}

func NewPostgresProductRepository(db *gorm.DB) *PostgresProductRepository {
	return &PostgresProductRepository{
		db: db,
	}
}

type PostgresWishListRepository struct {
	db *gorm.DB
}

func NewPostgresWishListRepository(db *gorm.DB) *PostgresWishListRepository {
	return &PostgresWishListRepository{
		db: db,
	}
}

type PostgresEnterpriseRepository struct {
	db *gorm.DB
}

func NewPostgresEnterpriseRepository(db *gorm.DB) *PostgresEnterpriseRepository {
	return &PostgresEnterpriseRepository{
		db: db,
	}
}

type PostgresRoleRepository struct {
	db *gorm.DB
}

func NewPostgresRoleRepository(db *gorm.DB) *PostgresRoleRepository {
	return &PostgresRoleRepository{
		db: db,
	}
}

type PostgresAuthRepository struct {
	db *gorm.DB
}

func NewPostgresAuthRepository(db *gorm.DB) *PostgresAuthRepository {
	return &PostgresAuthRepository{
		db: db,
	}
}
