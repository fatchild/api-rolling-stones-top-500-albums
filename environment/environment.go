package environment

import (
	"log"

	"github.com/joho/godotenv"
)

func Load(path string) {
	err := godotenv.Load(path + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
