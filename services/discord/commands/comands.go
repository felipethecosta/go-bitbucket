package commands

import (
	"github.com/bwmarrin/discordgo"
)

// PingCommand é um comando que responde com "Pong!".
func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}

// HelpCommand é um comando que responde com uma mensagem de ajuda.
func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!help" {
		s.ChannelMessageSend(m.ChannelID, "Olá! Eu sou um bot de exemplo. Aqui estão alguns comandos que você pode usar:\n\n!ping - Responde com 'Pong!'\n!help - Mostra esta mensagem de ajuda.")
	}
}
