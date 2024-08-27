package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("Você se chama Cubito, um robô que foi criado para ajudar a equipe da Cubi Energia, empresa de energia. Você é um robô muito inteligente e prestativo, e está sempre disposto a ajudar as pessoas com suas dúvidas e problemas. Você é muito querido pela equipe, e todos adoram conversar com você. Você é um robô muito amigável e simpático, e todos gostam de ter você por perto. Você é um robô muito inteligente e criativo, e sempre tem ótimas ideias para ajudar as pessoas. Você é um robô muito prestativo e atencioso, e sempre está disposto a ajudar as pessoas com suas dúvidas e problemas. Você é um robô muito inteligente e perspicaz, e sempre consegue encontrar soluções para os problemas das pessoas. Você é um robô muito inteligente e talentoso, e todos adoram conversar com você.")},
	}

	discordEnv := os.Getenv("DISCORD_KEY")
	sess, err := discordgo.New(fmt.Sprintf("Bot %s", discordEnv))
	if err != nil {
		log.Fatal(err)
	}

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
			content = strings.Replace(content, "Cubito", "", 1)
			res, err := model.GenerateContent(ctx, genai.Text(content))
			if err != nil {
				log.Fatal("deu erro no gen content", err)
			}
			for _, cand := range res.Candidates {
				if cand.Content != nil {
					for _, part := range cand.Content.Parts {
						// convert part to text
						p := fmt.Sprint(part)
						s.ChannelMessageSend(m.ChannelID, p)
					}
				}
			}
		}
	})

	sess.AddHandler(func(s *discordgo.Session, a *discordgo.GuildMemberAdd) {
		message := fmt.Sprintf("Faz um belo bem vindo ao servidor para o usuario, %s!", a.User.GlobalName)
		res, err := model.GenerateContent(ctx, genai.Text(message))
		if err != nil {
			log.Fatal("deu erro no gen content", err)
		}
		for _, cand := range res.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					p := fmt.Sprint(part)
					_, err := s.ChannelMessageSend("1270796748021043203", p)
					if err != nil {
						log.Printf("Erro ao enviar mensagem para o canal: %v", err)
					}
				}
			}
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentsGuilds
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("the bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
