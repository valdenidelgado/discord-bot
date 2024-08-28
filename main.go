package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/valdenidelgado/cubi-bot/ai"
	"github.com/valdenidelgado/cubi-bot/config"
	"github.com/valdenidelgado/cubi-bot/discord"
)

func main() {
	config.LoadEnv()

	session := discord.SetupDiscordBot()

	ctx := context.Background()
	client := ai.NewGenAIClient(ctx)

	discord.RegisterHandlers(session, client)

	defer session.Close()
	defer client.Client.Close()

	fmt.Println("the bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
