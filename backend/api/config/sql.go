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

func InitializeDbSQL(env *EnvVariables) (*gorm.DB, error) {
	urlDB := formatterUriDb("%s://%s:%s@%s/%s?sslmode=disable", env)

	db, err := gorm.Open(postgres.Open(urlDB), &gorm.Config{})
	if err != nil {
		loggerSQL.Errorf("Failed to connection to DataBase, error: %v", err)
		return nil, err
	}

	fmt.Println("The connection is successfully estabilize!")
	return db, nil
}

func TestConnectionSQL(database *gorm.DB) error {
	dbConn, err := database.DB()
	if err != nil {
		loggerSQL.Errorf("Expected valid sql.DB got error! Error: %v", err)
		return err
	}

	if err := dbConn.Ping(); err != nil {
		loggerSQL.Errorf("Expected to ping database successfully, got %v", err)
		return err
	}

	return nil
}
