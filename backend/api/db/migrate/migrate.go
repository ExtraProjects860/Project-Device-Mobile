package main

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"gorm.io/gorm"
)

func migrateDatabase(db *gorm.DB) error {
	var err error

	err = db.Migrator().DropTable(
		schemas.AllModelsSlice()...,
	)

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
