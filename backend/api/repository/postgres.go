package repository

import "gorm.io/gorm"

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository() *PostgresRepository {
	InitializeRepository()
	return &PostgresRepository{db: db}
}
