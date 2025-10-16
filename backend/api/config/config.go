package config

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*EnvVariables, *gorm.DB, error) {
	var err error

	env, err := InitilizeEnvVariables()
	if err != nil {
		return nil, nil, fmt.Errorf("error initializing env variables: %v", err)
	}

	db, err := InitializeDbServer(
		NewLogger("db"), 
		postgres.Open(env.DB.DbUri),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connection to DataBase, error: %v", err)
	}

	if err := db.AutoMigrate(
		schemas.AllModelsSlice()...,
	); err != nil {
		return nil, nil, fmt.Errorf("failed to migrate DB: %v", err)
	}

	return env, db, nil
}
