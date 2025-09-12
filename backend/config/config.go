package config

import (
	"log"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	env EnvVariables
)

func GetDB() *gorm.DB {
	return db
}

func GetEnv() EnvVariables {
	return env
}

func Init() error {
	var err error
	env, err = InitilizeEnvVariables()
	if err != nil {
		log.Fatalf("Failed to loading enviroment variables, error: %v", err)
	}

	db, err = InitializeDbSQL()
	if err != nil {
		log.Fatalf("Failed to connection to DataBase, error: %v", err)
	}

	return nil
}
