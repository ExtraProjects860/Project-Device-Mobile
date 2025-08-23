package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	urlDB := env.POSTGRESQL_URL

	db, err := gorm.Open(postgres.Open(urlDB), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connection to DataBase, error: %v", err)
	}
	fmt.Println("The connection is successfully estabilize!")

	return db, nil
}
