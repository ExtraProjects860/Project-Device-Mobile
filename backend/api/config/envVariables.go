package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Type     string
	Username string
	Password string
	Domain   string
	Name     string
}

type APIConfig struct {
	Port         string
	EmailService string
}

type EnvVariables struct {
	DB  DBConfig
	API APIConfig
}

func InitilizeEnvVariables() (EnvVariables, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	env := EnvVariables{
		DB: DBConfig{
			Type:     os.Getenv("DB_TYPE"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Domain:   os.Getenv("DB_DOMAIN"),
			Name:     os.Getenv("DB_NAME"),
		},
		API: APIConfig{
			Port:         os.Getenv("API_PORT"),
			EmailService: os.Getenv("EMAIL_SERVICE"),
		},
	}

	return env, nil
}
