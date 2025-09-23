package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/brianvoe/gofakeit/v7"
)

var (
	source                 = rand.NewPCG(uint64(time.Now().UnixNano()), rand.Uint64())
	rng    *rand.Rand      = rand.New(source)
	faker  *gofakeit.Faker = gofakeit.NewFaker(rng, true)
)

func verifyStartSeed(m any) bool {
	var count int64
	db.Model(m).Count(&count)
	return count > 0
}

func seedUser(quantity int) {
	modelName := "User"
	if verifyStartSeed(&schemas.User{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s' with %d records...", modelName, quantity)
	var role []schemas.Role
	var enterprise []schemas.Enterprise
	db.Find(&role)
	db.Find(&enterprise)

	hashedPassword, err := utils.GenerateHashPassword("123456")
	if err != nil {
		logger.Errorf("Error generate hash password. %v", err)
		panic(fmt.Sprintf("failed to hash password: %v", err))
	}

	for range quantity {
		re := role[rng.IntN(len(role))]
		en := enterprise[rng.IntN(len(enterprise))]

		user := schemas.User{
			Name:           faker.Name(),
			Email:          faker.Email(),
			Password:       hashedPassword,
			Cpf:            faker.Regex("[0-9]{11}"), // TODO pode gerar cpfs inválidos
			RegisterNumber: uint(faker.Number(1000, 9999)),
			RoleID:         re.ID,
			EnterpriseID:   &en.ID,
		}

		if err := db.CreateInBatches(&user, 5).Error; err != nil {
			logger.Errorf("Error creating %v: %v", modelName, err)
		}
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seedEnterprise(quantity int) {
	modelName := "Enterprise"
	if verifyStartSeed(&schemas.Enterprise{}) {
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

func seedRole() {
	modelName := "RoleUser"
	if verifyStartSeed(&schemas.Role{}) {
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

func seedWishList() {
	modelName := "WishList"
	if verifyStartSeed(&schemas.WishList{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	var users []schemas.User
	var products []schemas.Product
	db.Find(&users)
	db.Find(&products)

	if len(users) == 0 || len(products) == 0 {
		logger.Warningf("Cannot seed '%s'. Users or Products table is empty.", modelName)
		return
	}

	logger.Infof("Seeding table '%s'...", modelName)
	for range 10 {
		u := users[rand.IntN(len(users))]
		p := products[rand.IntN(len(products))]

		wishList := schemas.WishList{
			UserID:    u.ID,
			ProductID: p.ID,
		}

		if err := db.CreateInBatches(&wishList, 5).Error; err != nil {
			logger.Errorf("Error creating %v: %v", modelName, err)
		}
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seedProduct(quantity int) {
	modelName := "Product"
	if verifyStartSeed(&schemas.Product{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s' with %d records...", modelName, quantity)
	for range quantity {
		discount := faker.Price(5, 50) / 100

		product := schemas.Product{
			Name:               faker.ProductName(),
			Description:        faker.Sentence(10),
			Value:              faker.Price(50, 5000), // preço entre 50 e 5000
			Quantity:           faker.Number(1, 100),
			Discount:           &discount,
			IsPromotionAvaible: rng.IntN(2) == 1,
		}

		if err := db.CreateInBatches(&product, 5).Error; err != nil {
			logger.Errorf("Error creating %v: %v", modelName, err)
		}
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seeds() {
	seedRole()
	seedEnterprise(10)
	seedUser(30)
	seedProduct(30)
	seedWishList()

	logger.Info("Seed completed.")
}
