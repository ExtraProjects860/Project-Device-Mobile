package main

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/model"
	"gorm.io/gorm"
)

func migrateDatabase(db *gorm.DB) error {
	var err error

	err = db.Migrator().DropTable(
		model.AllModelsSlice()...,
	)

	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		model.AllModelsSlice()...,
	)

	if err != nil {
		return err
	}

	logger.Infof("Migrate models with successfully!")
	return nil
}
