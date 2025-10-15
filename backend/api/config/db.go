package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ExtraProjects860/Project-Device-Mobile/enum"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ResetDB(db *gorm.DB, logger *Logger) error {
	logger.Info("Dropping all tables...")
	err := db.Migrator().DropTable(
		schemas.AllModelsSlice()...,
	)
	if err != nil {
		return err
	}

	logger.Info("Recreating tables...")
	err = db.AutoMigrate(
		schemas.AllModelsSlice()...,
	)
	if err != nil {
		return err
	}

	logger.Info("Database reset completed.")
	return nil
}

func InitializeDbFile(logger *Logger, sqliteType string, path string) (*gorm.DB, error) {
	var dsn string

	switch sqliteType {
	case enum.Memory.String():
		dsn = ":memory"
	case enum.File.String():
		if path == "" {
			return nil, errors.New("sqlite path cannot be empty for file mode")
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return nil, err
		}

		dsn = path
	default:
		return nil, fmt.Errorf("unsupported sqlite type: %s", sqliteType)
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connection to DataBase, error: %v", err)
		return nil, err
	}

	logger.Info("Database connection successfully established!")
	return db, nil
}

func InitializeDbServer(logger *Logger, dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connection to DataBase, error: %v", err)
		return nil, err
	}

	logger.Info("Database connection successfully established!")
	return db, nil
}

func TestConnectionSQL(database *gorm.DB) error {
	logger := NewLogger("sql")

	dbConn, err := database.DB()
	if err != nil {
		logger.Errorf("Expected valid sql.DB got error! Error: %v", err)
		return err
	}

	if err := dbConn.Ping(); err != nil {
		logger.Errorf("Expected to ping database successfully, got %v", err)
		return err
	}

	return nil
}
