package config

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
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
