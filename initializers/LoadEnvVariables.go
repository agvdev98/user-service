package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
