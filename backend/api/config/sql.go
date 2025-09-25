package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var loggerSQL *Logger = GetLogger("sql")

func formatterUriDb(format string, env *EnvVariables) string {
	return fmt.Sprintf(
		format,
		env.DB.Type,
		env.DB.Username,
		env.DB.Password,
		env.DB.Domain,
		env.DB.Name,
	)
}

func InitializeDbSQL() (*gorm.DB, error) {
	urlDB := formatterUriDb("%s://%s:%s@%s/%s?sslmode=disable", GetEnv())

	db, err := gorm.Open(postgres.Open(urlDB), &gorm.Config{})
	if err != nil {
		loggerSQL.Errorf("Failed to connection to DataBase, error: %v", err)
		return nil, err
	}

	fmt.Println("The connection is successfully estabilize!")
	return db, nil
}

func TestConnectionSQL() error {
	db, err := GetDB().DB()
	if err != nil {
		loggerSQL.Errorf("Expected valid sql.DB got error! Error: %v", err)
		return err
	}

	if err := db.Ping(); err != nil {
		loggerSQL.Errorf("Expected to ping database successfully, got %v", err)
		return err
	}

	return nil
}
