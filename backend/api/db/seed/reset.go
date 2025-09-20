package main

import "github.com/ExtraProjects860/Project-Device-Mobile/model"

func resetDB() {
	logger.Info("Dropping all tables...")
	db.Migrator().DropTable(
		model.AllModelsSlice()...,
	)

	logger.Info("Recreating tables...")
	db.AutoMigrate(
		model.AllModelsSlice()...,
	)

	logger.Info("Database reset completed.")
}
