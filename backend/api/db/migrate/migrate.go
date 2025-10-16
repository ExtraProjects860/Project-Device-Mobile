package main

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"gorm.io/gorm"
)

func migrateDatabase(db *gorm.DB, logger *config.Logger) error {
	var err error

	err = config.ResetDB(db, logger)

	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		schemas.AllModelsSlice()...,
	)

	if err != nil {
		return err
	}

	logger.Infof("Migrate models with successfully!")
	return nil
}
