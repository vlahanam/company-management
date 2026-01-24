package initialize

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB    DB
	Fiber Fiber
	Auth  Auth
	CORS  CORS
}

type DB struct {
	DBPort     string
	DBHost     string
	DBUsername string
	DBPassword string
	DBName     string
}

type Fiber struct {
	Port string
}

type Auth struct {
	AccessSecret  string
	RefreshSecret string
}

type CORS struct {
	AllowedOrigins string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		DB: DB{
			DBPort:     os.Getenv("DB_PORT"),
			DBHost:     os.Getenv("DB_HOST"),
			DBUsername: os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
		},
		Fiber: Fiber{
			Port: os.Getenv("PORT"),
		},
		Auth: Auth{
			AccessSecret:  os.Getenv("ACCESS_SECRET_KEY"),
			RefreshSecret: os.Getenv("REFRESH_SECRET_KEY"),
		},
		CORS: CORS{
			AllowedOrigins: os.Getenv("CORS_ALLOWED_ORIGINS"),
		},
	}

	return cfg
}
