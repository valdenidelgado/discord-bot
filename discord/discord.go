package discord

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/valdenidelgado/cubi-bot/ai"
	"github.com/valdenidelgado/cubi-bot/cubi/api"
	"github.com/valdenidelgado/cubi-bot/data"
)

type Bot struct {
	Session *discordgo.Session
	Api     *api.API
}

func NewBot() *Bot {
	api := api.New()
	sess := setupDiscordSession()
	return &Bot{
		Session: sess,
		Api:     api,
	}
}

func setupDiscordSession() *discordgo.Session {
	sess, err := discordgo.New(fmt.Sprintf("Bot %s", data.DiscordKeyV2))
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions
	err = sess.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	return sess
}

func (b *Bot) RegisterHandlers(client *ai.GeminiAI) {
	b.RegisterCommands()

	b.Session.AddHandler(b.interactionCreate)
	b.Session.AddHandler(b.autocomplete)

	b.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
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
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Reviewers da sprint %s -> %s -> %s -> %s", reviewers[0], reviewers[1], reviewers[2], reviewers[3]))
		}

		if strings.Contains(content, "Cubito") || strings.Contains(content, "cubito") {
			responses := client.GenerateMessage(content)
			for _, response := range responses {
				s.ChannelMessageSend(m.ChannelID, response)
			}
		}
	})

	b.Session.AddHandler(func(s *discordgo.Session, a *discordgo.GuildMemberAdd) {
		message := fmt.Sprintf("Faz um belo bem vindo ao servidor para o usuario, %s!", a.User.GlobalName)
		channelID := "870419677434744865"
		responses := client.GenerateMessage(message)
		for _, response := range responses {
			_, err := s.ChannelMessageSend(channelID, response)
			if err != nil {
				log.Printf("Erro ao enviar mensagem para o canal: %v", err)
			}
		}
	})
}
