package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DbUri        string
	SqliteTypeDb string
	SqlitePath   string
}

type APIConfig struct {
	JwtKey       string
	RefreshKey   string
	Port         string
	EmailService string
}

type EnvVariables struct {
	DB  DBConfig
	API APIConfig
}

func InitilizeEnvVariables(paths ...string) (*EnvVariables, error) {
	if len(paths) == 0 {
		paths = []string{".env"}
	}

	loaded := false
	for _, path := range paths {
		if err := godotenv.Load(path); err == nil {
			loaded = true
			break
		}
	}

	if !loaded {
		return nil, fmt.Errorf("failed to load .env file from paths: %v", paths)
	}

	env := &EnvVariables{
		DB: DBConfig{
			DbUri:        os.Getenv("DB_URI"),
			SqliteTypeDb: os.Getenv("SQLITE_TYPE_DB"),
			SqlitePath:   os.Getenv("SQLITE_PATH"),
		},
		API: APIConfig{
			JwtKey:       os.Getenv("JWT_KEY"),
			RefreshKey:   os.Getenv("REFRESH_KEY"),
			Port:         os.Getenv("API_PORT"),
			EmailService: os.Getenv("EMAIL_SERVICE"),
		},
	}

	return env, nil
}
