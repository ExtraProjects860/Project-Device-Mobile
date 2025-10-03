package repository

import (
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
func NewPostgresUserRepository() UserRepository {
=======
func NewPostgresUserRepository(db *gorm.DB) UserRepository {
>>>>>>> dev
	return &postgresUserRepository{db: db}
}

type postgresProductRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
func NewPostgresProductRepository() ProductRepository {
=======
func NewPostgresProductRepository(db *gorm.DB) ProductRepository {
>>>>>>> dev
	return &postgresProductRepository{db: db}
}

type postgresWishListRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
func NewPostgresWishListRepository() WishListRepository {
=======
func NewPostgresWishListRepository(db *gorm.DB) WishListRepository {
>>>>>>> dev
	return &postgresWishListRepository{db: db}
}

type postgresEnterpriseRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
func NewPostgresEnterpriseRepository() EnterpriseRepository {
=======
func NewPostgresEnterpriseRepository(db *gorm.DB) EnterpriseRepository {
>>>>>>> dev
	return &postgresEnterpriseRepository{db: db}
}

type postgresRoleRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
func NewPostgresRoleRepository() RoleRepository {
=======
func NewPostgresRoleRepository(db *gorm.DB) RoleRepository {
>>>>>>> dev
	return &postgresRoleRepository{db: db}
}

type postgresTokenPasswordRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
func NewPostgresTokenPasswordRepository() TokenPasswordRepository {
=======
func NewPostgresTokenPasswordRepository(db *gorm.DB) TokenPasswordRepository {
>>>>>>> dev
	return &postgresTokenPasswordRepository{db: db}
}
