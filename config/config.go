package config

import (
	"fmt"
	"os"

)

func LoadEnv() {
	fmt.Print("Load commented\n")
	// err := godotenv.Load()
	// if err != nil {
		//log.Fatal("Error loading .env file")
	// }
}
func GetAPIKey() string {
	return os.Getenv("API_KEY")
}

func GetDiscordKey() string {
	return os.Getenv("DISCORD_KEY")
}

func GetChannelKey() string {
	return os.Getenv("DISCORD_CHANNEL_ID")
}
