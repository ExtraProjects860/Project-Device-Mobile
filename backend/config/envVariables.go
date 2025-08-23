package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	POSTGRESQL_URL string
}

func InitilizeEnvVariables() (EnvVariables, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	env := EnvVariables{
		os.Getenv("POSTGRESQL_URL"),
	}

	return env, nil
}
