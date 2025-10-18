package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/enum"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

var (
	source                 = rand.NewPCG(uint64(time.Now().UnixNano()), rand.Uint64())
	rng    *rand.Rand      = rand.New(source)
	faker  *gofakeit.Faker = gofakeit.NewFaker(rng, true)
)

func verifyStartSeed(db *gorm.DB, m any) bool {
	var count int64
	db.Model(m).Count(&count)
	return count > 0
}

func seedAdmin(db *gorm.DB, logger *config.Logger, enterprises []schemas.Enterprise) {
	hashedPassword, err := utils.GenerateHashPassword("admin")
	if err != nil {
		panic(fmt.Sprintf("failed to hash password: %v", err))
	}

	en := enterprises[rng.IntN(len(enterprises))]
	// url := utils.GenerateRandomPhoto(faker.Username(), 300, 300)

	user := schemas.User{
		Name:           strings.ToUpper("admin"),
		Email:          "admin@admin.com",
		Password:       hashedPassword,
		Cpf:            utils.GenerateCPF(),
		RegisterNumber: "1231231",
		PhotoUrl:       nil,
		RoleID:         1,
		EnterpriseID:   &en.ID,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("Error creating admin: %v", err)
	}
}

func seedUser(db *gorm.DB, logger *config.Logger, quantity int) {
	userModelName := "User"
	if verifyStartSeed(db, &schemas.User{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", userModelName)
		return
	}

	var roles []schemas.Role
	var enterprises []schemas.Enterprise
	db.Find(&roles)
	db.Find(&enterprises)

	if len(roles) == 0 || len(enterprises) == 0 {
		logger.Warningf("Cannot seed Users. Roles, Enterprises table is empty.")
		return
	}

	seedAdmin(db, logger, enterprises)

	logger.Infof("Seeding table '%s' with %d records...", userModelName, quantity)
	hashedPassword, err := utils.GenerateHashPassword("123456")
	if err != nil {
		panic(fmt.Sprintf("failed to hash password: %v", err))
	}

	usersToCreate := make([]schemas.User, 0, quantity)

	for range quantity {
		re := roles[rng.IntN(len(roles))]
		en := enterprises[rng.IntN(len(enterprises))]
		url := utils.GenerateRandomPhoto(faker.Username(), 300, 300)

		user := schemas.User{
			Name:           strings.ToUpper(faker.Name()),
			Email:          faker.Email(),
			Password:       hashedPassword,
			Cpf:            utils.GenerateCPF(),
			RegisterNumber: strconv.Itoa(faker.Number(1000000, 9999999)),
			PhotoUrl:       &url,
			RoleID:         re.ID,
			EnterpriseID:   &en.ID,
		}
		usersToCreate = append(usersToCreate, user)
	}

	if err := db.CreateInBatches(usersToCreate, 10).Error; err != nil {
		logger.Errorf("Error creating %s: %v", userModelName, err)
	}

	logger.Infof("Seeding for table '%s' completed.", userModelName)
}

func seedEnterprise(db *gorm.DB, logger *config.Logger, quantity int) {
	modelName := "Enterprise"
	if verifyStartSeed(db, &schemas.Enterprise{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s'...", modelName)
	for range quantity {
		enterprise := schemas.Enterprise{
			Name: strings.ToUpper(faker.AppName()),
		}

		if err := db.CreateInBatches(&enterprise, 3).Error; err != nil {
			logger.Errorf("Error creating %v: %v", modelName, err)
		}
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seedWishList(db *gorm.DB, logger *config.Logger) {
	wishlistModelName := "WishList"
	if verifyStartSeed(db, &schemas.WishList{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", wishlistModelName)
		return
	}

	logger.Infof("Seeding data for '%s'...", wishlistModelName)

	var users []schemas.User
	var products []schemas.Product
	db.Find(&users)
	db.Find(&products)

	if len(users) == 0 || len(products) == 0 {
		logger.Warningf("Cannot seed '%s'. Users or Products table is empty.", wishlistModelName)
		return
	}

	wishlistEntriesToCreate := make([]schemas.WishList, 0)
	for _, user := range users {
		numProductsInWishlist := rng.IntN(6) + 2

		productsUsedInThisWishlist := make(map[uint]bool)

		for range numProductsInWishlist {
			randomProduct := products[rng.IntN(len(products))]

			if productsUsedInThisWishlist[randomProduct.ID] {
				continue
			}

			productsUsedInThisWishlist[randomProduct.ID] = true

			entry := schemas.WishList{
				UserID:    user.ID,
				ProductID: randomProduct.ID,
			}

			wishlistEntriesToCreate = append(wishlistEntriesToCreate, entry)
		}
	}

	if len(wishlistEntriesToCreate) > 0 {
		if err := db.CreateInBatches(wishlistEntriesToCreate, 100).Error; err != nil {
			logger.Errorf("Error creating %s entries: %v", wishlistModelName, err)
		}
	}

	logger.Infof("Seeding for '%s' completed.", wishlistModelName)
}

func seedRole(db *gorm.DB, logger *config.Logger) {
	modelName := "RoleUser"
	if verifyStartSeed(db, &schemas.Role{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s'...", modelName)
	roles := []schemas.Role{
		{Name: enum.SuperAdmin.String()},
		{Name: enum.Admin.String()},
		{Name: enum.User.String()},
	}

	if err := db.Create(&roles).Error; err != nil {
		logger.Errorf("Error creating %v: %v", modelName, err)
	}
	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seedProduct(db *gorm.DB, logger *config.Logger, quantity int) {
	modelName := "Product"
	if verifyStartSeed(db, &schemas.Product{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s' with %d records...", modelName, quantity)
	for range quantity {
		url := utils.GenerateRandomPhoto(faker.Username(), 300, 300)
		discount := faker.Price(5, 50) / 100
		randAvaiable := rng.IntN(2) == 1

		product := schemas.Product{
			Name:               strings.ToUpper(faker.ProductName()),
			Description:        faker.Sentence(10),
			Value:              faker.Price(50, 5000), // preço entre 50 e 5000
			Quantity:           faker.Number(1, 100),
			Discount:           &discount,
			PhotoUrl:           &url,
			IsPromotionAvaible: &randAvaiable,
		}

		if err := db.CreateInBatches(&product, 5).Error; err != nil {
			logger.Errorf("Error creating %v: %v", modelName, err)
		}
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seeds(db *gorm.DB, logger *config.Logger) {
	seedRole(db, logger)
	seedEnterprise(db, logger, 10)
	seedProduct(db, logger, 30)
	seedUser(db, logger, 30)
	seedWishList(db, logger)

	logger.Info("Seed completed.")
}
