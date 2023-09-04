package database

import (
	"os"

	"github.com/joho/godotenv"
)

func EnvDatabase() Config {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	return Config{
		CONSUMER_KEY:    os.Getenv("CONSUMER_KEY"),
		CONSUMER_SECRET: os.Getenv("CONSUMER_SECRET"),
	}
}
