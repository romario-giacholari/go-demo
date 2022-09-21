package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() map[string]string {
	envs, err := godotenv.Read()

	if err != nil {
		log.Fatal("Failed to load the .env file")
	}

	return envs
}
