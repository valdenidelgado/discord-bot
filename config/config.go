package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetAPIKey() string {
	return os.Getenv("API_KEY")
}

func GetDiscordKey() string {
	return os.Getenv("DISCORD_KEY")
}
