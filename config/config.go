package config

import (
	"os"
)

func GetAPIKey() string {
	return os.Getenv("API_KEY")
}

func GetDiscordKey() string {
	return os.Getenv("DISCORD_KEY")
}

func GetChannelKey() string {
	return os.Getenv("DISCORD_CHANNEL_ID")
}
