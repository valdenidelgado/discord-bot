package ai

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/valdenidelgado/cubi-bot/config"
	"google.golang.org/api/option"
)

type GeminiAI struct {
	Client *genai.Client
	model  *genai.GenerativeModel
	ctx    context.Context
}

func NewGenAIClient(ctx context.Context) *GeminiAI {
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GetAPIKey()))
	if err != nil {
		log.Fatal("client error", err)
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("Você se chama Cubito, um robô que foi criado para ajudar a equipe da Cubi Energia, empresa de energia. Você é um robô muito inteligente e prestativo, e está sempre disposto a ajudar as pessoas com suas dúvidas e problemas. Você é muito querido pela equipe, e todos adoram conversar com você. Você é um robô muito amigável e simpático, e todos gostam de ter você por perto. Você é um robô muito inteligente e criativo, e sempre tem ótimas ideias para ajudar as pessoas. Você é um robô muito prestativo e atencioso, e sempre está disposto a ajudar as pessoas com suas dúvidas e problemas. Você é um robô muito inteligente e perspicaz, e sempre consegue encontrar soluções para os problemas das pessoas. Você é um robô muito inteligente e talentoso, e todos adoram conversar com você.")},
	}

	return &GeminiAI{
		Client: client,
		model:  model,
		ctx:    ctx,
	}
}

func (g *GeminiAI) GenerateMessage(message string) []string {
	res, err := g.model.GenerateContent(g.ctx, genai.Text(message))
	if err != nil {
		log.Fatal("Error generating content", err)
	}

	var responses []string
	for _, cand := range res.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				responses = append(responses, fmt.Sprintf("%s", part))
			}
		}
	}

	return responses
}
