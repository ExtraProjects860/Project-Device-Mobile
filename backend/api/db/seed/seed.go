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

// TODO colocar concorrência, porém pra cada função separadamente

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

	const password string = "123456"

	for i := 0; i < quantity; i++ {
		re := role[rng.IntN(len(role))]
		en := enterprise[rng.IntN(len(enterprise))]

		hashedPassword, err := utils.GenerateHashPassword(password)
		if err != nil {
			logger.Errorf("Error generate hash password. %v", err)
			panic(fmt.Sprintf("failed to hash password: %v", err))
		}

		user := schemas.User{
			Name:           faker.Name(),
			Email:          faker.Email(),
			Password:       hashedPassword,           // TODO essa senha dps vai ser substituída pela função hash
			Cpf:            faker.Regex("[0-9]{11}"), // TODO pode gerar cpfs inválidos
			RegisterNumber: uint(faker.Number(1000, 9999)),
			RoleID:         re.ID,
			EnterpriseID:   &en.ID,
		}

		db.Create(&user)
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
	for i := 0; i < quantity; i++ {
		enterprise := schemas.Enterprise{
			Name: faker.AppName(),
		}

		db.Create(&enterprise)
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

	db.Create(&roles)
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
	for i := 0; i < 10; i++ {
		u := users[rand.IntN(len(users))]
		p := products[rand.IntN(len(products))]

		w := schemas.WishList{
			UserID:    u.ID,
			ProductID: p.ID,
		}

		db.Create(&w)
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
	for i := 0; i < quantity; i++ {
		discount := faker.Price(5, 50) / 100

		product := schemas.Product{
			Name:               faker.ProductName(),
			Description:        faker.Sentence(10),
			Value:              faker.Price(50, 5000), // preço entre 50 e 5000
			Quantity:           faker.Number(1, 100),
			Discount:           &discount,
			IsPromotionAvaible: rng.IntN(2) == 1,
		}
		db.Create(&product)
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
