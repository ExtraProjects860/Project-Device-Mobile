package repository

import (
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) UserRepository {
	return &postgresUserRepository{db: db}
}

type postgresProductRepository struct {
	db *gorm.DB
}

func NewPostgresProductRepository(db *gorm.DB) ProductRepository {
	return &postgresProductRepository{db: db}
}

type postgresWishListRepository struct {
	db *gorm.DB
}

func NewPostgresWishListRepository(db *gorm.DB) WishListRepository {
	return &postgresWishListRepository{db: db}
}

type postgresEnterpriseRepository struct {
	db *gorm.DB
}

func NewPostgresEnterpriseRepository(db *gorm.DB) EnterpriseRepository {
	return &postgresEnterpriseRepository{db: db}
}

type postgresRoleRepository struct {
	db *gorm.DB
}

func NewPostgresRoleRepository(db *gorm.DB) RoleRepository {
	return &postgresRoleRepository{db: db}
}

type postgresTokenPasswordRepository struct {
	db *gorm.DB
}

func NewPostgresTokenPasswordRepository(db *gorm.DB) TokenPasswordRepository {
	return &postgresTokenPasswordRepository{db: db}
}
