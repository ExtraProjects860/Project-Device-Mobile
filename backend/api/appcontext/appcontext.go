package appcontext

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type AppContext struct {
	Env        *config.EnvVariables
	Cloudinary *cloudinary.Cloudinary
	DB         *gorm.DB
}

func SetupContext(env *config.EnvVariables, db *gorm.DB) (*AppContext, error) {
	cld, err := cloudinary.NewFromURL(env.API.Cloudnary)
	if err != nil {
		return nil, err
	}

	return &AppContext{
		Env: env,
		DB:  db,
		Cloudinary: cld,
	}, nil
}
