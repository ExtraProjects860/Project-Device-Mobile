package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
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

<<<<<<< HEAD
func seedUser(quantity int) {
	userModelName := "User"
	if verifyStartSeed(&schemas.User{}) {
=======
func seedUser(db *gorm.DB, quantity int) {
	userModelName := "User"
	if verifyStartSeed(db, &schemas.User{}) {
>>>>>>> dev
		logger.Infof("Table '%s' already has data. Skipping seed.", userModelName)
		return
	}

	var roles []schemas.Role
	var enterprises []schemas.Enterprise
	var products []schemas.Product
	db.Find(&roles)
	db.Find(&enterprises)
	db.Find(&products)

	if len(roles) == 0 || len(enterprises) == 0 || len(products) == 0 {
		logger.Warningf("Cannot seed Users. Roles, Enterprises or Products table is empty.")
		return
	}

	logger.Infof("Seeding table '%s' with %d records...", userModelName, quantity)
	hashedPassword, err := utils.GenerateHashPassword("123456")
	if err != nil {
		panic(fmt.Sprintf("failed to hash password: %v", err))
	}

	usersToCreate := make([]schemas.User, 0, quantity)

	for i := 0; i < quantity; i++ {
		re := roles[rng.IntN(len(roles))]
		en := enterprises[rng.IntN(len(enterprises))]
		url := fmt.Sprintf("https://picsum.photos/%d/%d?random=%s", 300, 300, faker.Username())

		user := schemas.User{
			Name:           faker.Name(),
			Email:          faker.Email(),
			Password:       hashedPassword,
			Cpf:            faker.Regex("[0-9]{11}"),
			RegisterNumber: uint(faker.Number(1000, 9999)),
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

func seedEnterprise(db *gorm.DB, quantity int) {
	modelName := "Enterprise"
	if verifyStartSeed(db, &schemas.Enterprise{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s'...", modelName)
	for range quantity {
		enterprise := schemas.Enterprise{
			Name: faker.AppName(),
		}

		if err := db.CreateInBatches(&enterprise, 3).Error; err != nil {
			logger.Errorf("Error creating %v: %v", modelName, err)
		}
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

<<<<<<< HEAD
func seedWishList() {
	wishlistModelName := "WishList"
	if verifyStartSeed(&schemas.WishList{}) {
=======
func seedWishList(db *gorm.DB) {
	wishlistModelName := "WishList"
	if verifyStartSeed(db, &schemas.WishList{}) {
>>>>>>> dev
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

<<<<<<< HEAD
func seedRole() {
=======
func seedRole(db *gorm.DB) {
>>>>>>> dev
	modelName := "RoleUser"
	if verifyStartSeed(db, &schemas.Role{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s'...", modelName)
	roles := []schemas.Role{
		{Name: config.SuperAdmin.String()},
		{Name: config.Admin.String()},
		{Name: config.User.String()},
	}

	if err := db.Create(&roles).Error; err != nil {
		logger.Errorf("Error creating %v: %v", modelName, err)
	}
	logger.Infof("Seeding for table '%s' completed.", modelName)
}

<<<<<<< HEAD
func seedProduct(quantity int) {
=======
func seedProduct(db *gorm.DB, quantity int) {
>>>>>>> dev
	modelName := "Product"
	if verifyStartSeed(db, &schemas.Product{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s' with %d records...", modelName, quantity)
	for range quantity {
		url := fmt.Sprintf("https://picsum.photos/%d/%d?random=%s", 300, 300, faker.Username())
		discount := faker.Price(5, 50) / 100

		product := schemas.Product{
			Name:               faker.ProductName(),
			Description:        faker.Sentence(10),
			Value:              faker.Price(50, 5000), // preço entre 50 e 5000
			Quantity:           faker.Number(1, 100),
			Discount:           &discount,
			PhotoUrl:           &url,
			IsPromotionAvaible: rng.IntN(2) == 1,
		}

		if err := db.CreateInBatches(&product, 5).Error; err != nil {
			logger.Errorf("Error creating %v: %v", modelName, err)
		}
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

<<<<<<< HEAD
func seeds() {
	seedRole()
	seedEnterprise(10)
	seedProduct(30)
	seedUser(30)
	seedWishList()
=======
func seeds(db *gorm.DB) {
	seedRole(db)
	seedEnterprise(db, 10)
	seedProduct(db, 30)
	seedUser(db, 30)
	seedWishList(db, )
>>>>>>> dev

	logger.Info("Seed completed.")
}
