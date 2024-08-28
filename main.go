package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/valdenidelgado/cubi-bot/ai"
	"github.com/valdenidelgado/cubi-bot/discord"
)

func main() {
	session := discord.SetupDiscordBot()

	ctx := context.Background()
	client := ai.NewGenAIClient(ctx)

	discord.RegisterHandlers(session, client)

	defer session.Close()
	defer client.Client.Close()

	fmt.Println("the bot is online")

	http.ListenAndServe(":8080", nil)
}
