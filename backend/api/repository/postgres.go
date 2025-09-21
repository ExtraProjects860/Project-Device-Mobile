package repository

import "gorm.io/gorm"

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository() *PostgresUserRepository {
	InitializeRepository()
	return &PostgresUserRepository{db: db}
}
