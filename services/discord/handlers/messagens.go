package handlers

import (
	"github.com/bwmarrin/discordgo"
)

// SendEmbedMessage envia uma mensagem com embed.
func SendEmbedMessage(s *discordgo.Session, channelID string) {
	embed := &discordgo.MessageEmbed{
		Title:       "Título do Embed",
		Description: "Esta é uma descrição do embed.",
		Color:       0x00ff00, // Cor verde
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Campo 1",
				Value:  "Valor do Campo 1",
				Inline: true,
			},
			{
				Name:   "Campo 2",
				Value:  "Valor do Campo 2",
				Inline: true,
			},
		},
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		s.ChannelMessageSend(channelID, "Erro ao enviar a mensagem embed.")
	}
}

// AddReaction adiciona uma reação a uma mensagem.
func AddReaction(s *discordgo.Session, channelID, messageID string, emoji string) {
	err := s.MessageReactionAdd(channelID, messageID, emoji)
	if err != nil {
		s.ChannelMessageSend(channelID, "Erro ao adicionar a reação.")
	}
}
