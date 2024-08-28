package discord

import (
	"fmt"
	"log"
	"strings"
	"time"

	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/valdenidelgado/cubi-bot/ai"
	"github.com/valdenidelgado/cubi-bot/config"
)

func SetupDiscordBot() *discordgo.Session {
	sess, err := discordgo.New(fmt.Sprintf("Bot %s", config.GetDiscordKey()))
	if err != nil {
		log.Fatal(err)
	}

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentsGuilds
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}

	return sess
}

func RegisterHandlers(sess *discordgo.Session, client *ai.GeminiAI) {

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		content := m.Content
		reviwers := "reviewers"

		if strings.HasPrefix(strings.ToLower(content), strings.ToLower(reviwers)) {
			reviewers := strings.Split(content, " ")
			reviewers = reviewers[1:]
			rand.New(rand.NewSource(time.Now().UnixNano()))
			rand.Shuffle(len(reviewers), func(i, j int) { reviewers[i], reviewers[j] = reviewers[j], reviewers[i] })
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Reviewers da sprint %s -> %s -> %s", reviewers[0], reviewers[1], reviewers[2]))
		}

		if strings.HasPrefix(content, "Cubito") || strings.HasPrefix(content, "cubito") {
			responses := client.GenerateMessage(content)
			for _, response := range responses {
				s.ChannelMessageSend(m.ChannelID, response)
			}
		}
	})

	sess.AddHandler(func(s *discordgo.Session, a *discordgo.GuildMemberAdd) {
		message := fmt.Sprintf("Faz um belo bem vindo ao servidor para o usuario, %s!", a.User.GlobalName)
		// channelID := os.Getenv("DISCORD_CHANNEL_ID")
		channelID := config.GetChannelKey()
		responses := client.GenerateMessage(message)
		for _, response := range responses {
			_, err := s.ChannelMessageSend(channelID, response)
			if err != nil {
				log.Printf("Erro ao enviar mensagem para o canal: %v", err)
			}
		}
	})
}
