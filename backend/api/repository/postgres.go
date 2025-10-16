package repository

import (
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
<<<<<<< HEAD
func NewPostgresUserRepository() UserRepository {
=======
func NewPostgresUserRepository(db *gorm.DB) UserRepository {
>>>>>>> dev
	return &postgresUserRepository{db: db}
=======
func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
>>>>>>> dev
}

type PostgresProductRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
<<<<<<< HEAD
func NewPostgresProductRepository() ProductRepository {
=======
func NewPostgresProductRepository(db *gorm.DB) ProductRepository {
>>>>>>> dev
	return &postgresProductRepository{db: db}
=======
func NewPostgresProductRepository(db *gorm.DB) *PostgresProductRepository {
	return &PostgresProductRepository{
		db: db,
	}
>>>>>>> dev
}

type PostgresWishListRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
<<<<<<< HEAD
func NewPostgresWishListRepository() WishListRepository {
=======
func NewPostgresWishListRepository(db *gorm.DB) WishListRepository {
>>>>>>> dev
	return &postgresWishListRepository{db: db}
=======
func NewPostgresWishListRepository(db *gorm.DB) *PostgresWishListRepository {
	return &PostgresWishListRepository{
		db: db,
	}
>>>>>>> dev
}

type PostgresEnterpriseRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
<<<<<<< HEAD
func NewPostgresEnterpriseRepository() EnterpriseRepository {
=======
func NewPostgresEnterpriseRepository(db *gorm.DB) EnterpriseRepository {
>>>>>>> dev
	return &postgresEnterpriseRepository{db: db}
=======
func NewPostgresEnterpriseRepository(db *gorm.DB) *PostgresEnterpriseRepository {
	return &PostgresEnterpriseRepository{
		db: db,
	}
>>>>>>> dev
}

type PostgresRoleRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
<<<<<<< HEAD
func NewPostgresRoleRepository() RoleRepository {
=======
func NewPostgresRoleRepository(db *gorm.DB) RoleRepository {
>>>>>>> dev
	return &postgresRoleRepository{db: db}
=======
func NewPostgresRoleRepository(db *gorm.DB) *PostgresRoleRepository {
	return &PostgresRoleRepository{
		db: db,
	}
>>>>>>> dev
}

type PostgresAuthRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
<<<<<<< HEAD
func NewPostgresTokenPasswordRepository() TokenPasswordRepository {
=======
func NewPostgresTokenPasswordRepository(db *gorm.DB) TokenPasswordRepository {
>>>>>>> dev
	return &postgresTokenPasswordRepository{db: db}
=======
func NewPostgresAuthRepository(db *gorm.DB) *PostgresAuthRepository {
	return &PostgresAuthRepository{
		db: db,
	}
>>>>>>> dev
}
