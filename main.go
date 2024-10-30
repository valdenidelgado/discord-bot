package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/valdenidelgado/cubi-bot/ai"
	"github.com/valdenidelgado/cubi-bot/config"
	"github.com/valdenidelgado/cubi-bot/discord"
)

func main() {
	defer handlePanic()
	config.LoadEnv()

	sess := discord.NewBot()

	ctx := context.Background()
	client := ai.NewGenAIClient(ctx)

	sess.RegisterHandlers(client)

	defer sess.Session.Close()
	defer client.Client.Close()

	fmt.Println("the bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func handlePanic() {
	if r := recover(); r != nil {
		log.Printf("Recuperado de um erro crítico: %v", r)
	}
}
