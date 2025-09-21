package main

import "github.com/ExtraProjects860/Project-Device-Mobile/schemas"

func resetDB() {
	logger.Info("Dropping all tables...")
	db.Migrator().DropTable(
		schemas.AllModelsSlice()...,
	)

	logger.Info("Recreating tables...")
	db.AutoMigrate(
		schemas.AllModelsSlice()...,
	)

	logger.Info("Database reset completed.")
}
