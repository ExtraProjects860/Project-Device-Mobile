package config

import (
	"log"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

func Init() error {
	var err error
	db, err = InitializePostgreSQL()
	if err != nil {
		log.Fatalf("Failed to connection to DataBase, error: %v", err)
	}
	return nil
}
