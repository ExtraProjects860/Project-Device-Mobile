package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	env    EnvVariables
)

func GetDB() *gorm.DB {
	return db
}

func GetEnv() EnvVariables {
	return env
}

func GetLogger(prefix string) *Logger {
	return NewLogger(prefix)
}

func Init() error {
	var err error

	env, err = InitilizeEnvVariables()
	if err != nil {
		return fmt.Errorf("error initializing env variables: %v", err)
	}

	db, err = InitializeDbSQL()
	if err != nil {
		return fmt.Errorf("failed to connection to DataBase, error: %v", err)
	}

	return nil
}
