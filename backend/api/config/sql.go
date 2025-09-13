package config

import (
	"fmt"

	"github.com/ExtraProjects860/Project-Device-Mobile/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var loggerSQL = GetLogger("sql")

func formatterUriDb(format string, env EnvVariables) string {
	return fmt.Sprintf(
		format,
		env.DB.Type,
		env.DB.Username,
		env.DB.Password,
		env.DB.Domain,
		env.DB.Name,
	)
}

func migrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.TypeUser{},
		&model.User{},
		&model.TokenPassword{},
		&model.WishList{},
		&model.Product{},
		&model.Promotion{},
	)
	if err != nil {
		return err
	}
	return nil
}

func InitializeDbSQL() (*gorm.DB, error) {
	urlDB := formatterUriDb("%s://%s:%s@%s/%s?sslmode=disable", GetEnv())

	db, err := gorm.Open(postgres.Open(urlDB), &gorm.Config{})
	if err != nil {
		loggerSQL.Errorf("Failed to connection to DataBase, error: %v", err)
		return nil, err
	}

	err = migrateDatabase(db)
	if err != nil {
		loggerSQL.Errorf("AutoMigrate error: %v", err)
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
	defer db.Close()

	if err := db.Ping(); err != nil {
		loggerSQL.Errorf("Expected to ping database successfully, got %v", err)
		return err
	}

	return nil
}
