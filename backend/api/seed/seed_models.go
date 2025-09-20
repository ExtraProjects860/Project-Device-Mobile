package seed

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/model"
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
	if verifyStartSeed(&model.User{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s' with %d records...", modelName, quantity)
	var typeUsers []model.TypeUser
	db.Find(&typeUsers)

	const password string = "123456"

	for i := 0; i < quantity; i++ {
		tu := typeUsers[rng.IntN(len(typeUsers))]

		hashedPassword, err := utils.GenerateHashPassword(password)
		if err != nil {
			logger.Errorf("Error generate hash password. %v", err)
			panic(fmt.Sprintf("failed to hash password: %v", err))
		}

		user := model.User{
			Name:           faker.Name(),
			Email:          faker.Email(),
			Password:       hashedPassword,                 // TODO essa senha dps vai ser substituída pela função hash
			Cpf:            faker.Regex("[0-9]{11}"), // TODO pode gerar cpfs inválidos
			RegisterNumber: uint(faker.Number(1000, 9999)),
			TypeUserID:     tu.ID,
		}

		db.Create(&user)
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seedTypeUser() {
	modelName := "TypeUser"
	if verifyStartSeed(&model.TypeUser{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s'...", modelName)
	types := []model.TypeUser{
		{Name: config.SuperAdmin.String()},
		{Name: config.Admin.String()},
		{Name: config.User.String()},
	}

	db.Create(&types)
	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seedWishList() {
	modelName := "WishList"
	if verifyStartSeed(&model.WishList{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	var users []model.User
	var products []model.Product
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

		w := model.WishList{
			UserID:    u.ID,
			ProductID: p.ID,
		}

		db.Create(&w)
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seedProduct(quantity int) {
	modelName := "Product"
	if verifyStartSeed(&model.Product{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	logger.Infof("Seeding table '%s' with %d records...", modelName, quantity)
	for i := 0; i < quantity; i++ {
		product := model.Product{
			Name:        faker.ProductName(),
			Description: faker.Sentence(10),
			Value:       faker.Price(50, 5000), // preço entre 50 e 5000
			Quantity:    faker.Number(1, 100),
		}
		db.Create(&product)
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func seedPromotion() {
	modelName := "Promotion"
	if verifyStartSeed(&model.Promotion{}) {
		logger.Infof("Table '%s' already has data. Skipping seed.", modelName)
		return
	}

	var products []model.Product
	db.Find(&products)

	if len(products) == 0 {
		logger.Warningf("Cannot seed '%s'. Products table is empty.", modelName)
		return
	}

	logger.Infof("Seeding table '%s'...", modelName)
	for i := 0; i < 5; i++ {
		p := products[rng.IntN(len(products))]
		discount := faker.Price(5, 50) / 100
		promotion := model.Promotion{
			ProductID: p.ID,
			Discount:  &discount,
		}
		db.Create(&promotion)
	}

	logger.Infof("Seeding for table '%s' completed.", modelName)
}

func Seeds() {
	seedTypeUser()
	seedUser(30)
	seedProduct(30)
	seedWishList()
	seedPromotion()

	logger.Info("Seed completed.")
}
