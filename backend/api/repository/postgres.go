package repository

import (
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository() UserRepository {
	return &postgresUserRepository{db: db}
}

type postgresProductRepository struct {
	db *gorm.DB
}

func NewPostgresProductRepository() ProductRepository {
	return &postgresProductRepository{db: db}
}

type postgresWishListRepository struct {
	db *gorm.DB
}

func NewPostgresWishListRepository() WishListRepository {
	return &postgresWishListRepository{db: db}
}

type postgresEnterpriseRepository struct {
	db *gorm.DB
}

func NewPostgresEnterpriseRepository() EnterpriseRepository {
	return &postgresEnterpriseRepository{db: db}
}

type postgresRoleRepository struct {
	db *gorm.DB
}

func NewPostgresRoleRepository() RoleRepository {
	return &postgresRoleRepository{db: db}
}

type postgresTokenPasswordRepository struct {
	db *gorm.DB
}

func NewPostgresTokenPasswordRepository() TokenPasswordRepository {
	return &postgresTokenPasswordRepository{db: db}
}
