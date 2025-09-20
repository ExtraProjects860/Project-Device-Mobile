package seed

import "github.com/ExtraProjects860/Project-Device-Mobile/model"

func ResetDB() {
	logger.Info("Dropping all tables...")
	db.Migrator().DropTable(
		&model.User{},
		&model.TypeUser{},
		&model.WishList{},
		&model.Product{},
		&model.Promotion{},
	)

	logger.Info("Recreating tables...")
	db.AutoMigrate(
		&model.User{},
		&model.TypeUser{},
		&model.WishList{},
		&model.Product{},
		&model.Promotion{},
	)

	logger.Info("Database reset completed.")
}
