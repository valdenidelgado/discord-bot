package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/valdenidelgado/cubi-bot/data"
)

func (b *Bot) RegisterCommands() {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "buscar",
			Description: "Buscar empresas, unidades e documentos",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "tipo",
					Description: "Tipo de busca",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: "Empresas", Value: "empresas"},
						{Name: "Unidades", Value: "unidades"},
						{Name: "Documento", Value: "documento"},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "criterio",
					Description: "Critério de busca",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: "ID", Value: "id"},
						{Name: "Nome", Value: "nome"},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "ID para buscar",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "nome",
					Description: "Nome para buscar",
					Required:    false,
				},
			},
		},
	}

	_, err := b.Session.ApplicationCommandBulkOverwrite("1278732626068508682", "", commands)
	if err != nil {
		log.Fatalf("Cannot create slash commands: %v", err)
	}
}

func (b *Bot) interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		options := i.ApplicationCommandData().Options
		tipo := options[0].StringValue()
		criterio := options[1].StringValue()
		var response string

		if tipo == "empresas" {
			if criterio == "id" {
				if len(options) < 3 {
					response = data.InfoId
				} else {
					id := options[2].StringValue()
					response = b.Api.GetCompanyById(id)
				}
			} else if criterio == "nome" {
				// nome := options[3].StringValue()
				response = data.PremiumMessage
			}
		}
		if tipo == "unidades" {
			if criterio == "id" {
				if len(options) < 3 {
					response = data.InfoId
				} else {
					id := options[2].StringValue()
					response = b.Api.GetBranchById(id)
				}
			} else if criterio == "nome" {
				// nome := options[3].StringValue()
				response = data.PremiumMessage
			}
		}
		if tipo == "documento" {
			if criterio == "id" {
				if len(options) < 3 {
					response = data.InfoId
				} else {
					id := options[2].StringValue()
					response = b.Api.GetBillingDetailsById(id)
				}
			} else if criterio == "nome" {
				// nome := options[3].StringValue()
				response = data.PremiumMessage
			}
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		})
		if err != nil {
			log.Printf("Erro ao responder à interação do comando: %v", err)
		}
	}
}

func (b *Bot) autocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommandAutocomplete {
		data := i.ApplicationCommandData()
		if data.Name == "buscar" {
			var choices []*discordgo.ApplicationCommandOptionChoice
			if data.Options[0].Name == "tipo" {
				choices = []*discordgo.ApplicationCommandOptionChoice{
					{Name: "Empresas", Value: "empresas"},
					{Name: "Unidades", Value: "unidades"},
					{Name: "Usuários", Value: "usuarios"},
				}
			} else if data.Options[0].Name == "criterio" {
				choices = []*discordgo.ApplicationCommandOptionChoice{
					{Name: "ID", Value: "id"},
					{Name: "Nome", Value: "nome"},
				}
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionApplicationCommandAutocompleteResult,
				Data: &discordgo.InteractionResponseData{
					Choices: choices,
				},
			})
			if err != nil {
				log.Printf("Erro ao responder à interação de autocompletar: %v", err)
			}
		}
	}
}
