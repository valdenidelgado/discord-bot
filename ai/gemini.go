package ai

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/valdenidelgado/cubi-bot/data"
	"google.golang.org/api/option"
)

type GeminiAI struct {
	Client *genai.Client
	model  *genai.GenerativeModel
	ctx    context.Context
}

func NewGenAIClient(ctx context.Context) *GeminiAI {
	client, err := genai.NewClient(ctx, option.WithAPIKey(data.GeminiKey))
	if err != nil {
		log.Fatal("client error", err)
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(data.PrePrompt)},
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
