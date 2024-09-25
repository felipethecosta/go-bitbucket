package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Variables used for command line parameters
var prefix = "!"

// Função principal para iniciar o bot
func main() {
    // Criação da sessão do Discord
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Variables de ambiente para o token do bot
    token := os.Getenv("DISCORD_TOKEN")
    channelID := os.Getenv("CHANNEL_ID")

    // Criação da sessão do Discord com o token do bot
    dg, err := discordgo.New("Bot " + token)
    if err != nil {
        fmt.Println("Error creating Discord session: ", err)
        return
    }

    // Envia uma mensagem para o canal espeficicado ao iniciar o bot
    _, _ = dg.ChannelMessageSend(channelID, "Bot está online, observe nosso prefix = ! e saiba mais.")
    if err != nil {
        fmt.Println("Error sending message: ", err)
        return
    }

    // Vamos criar uma func para quando usuário enviar !pr ele abrir um embed com botões em baixo de confirmar e negar.
           // Criação da sessão do Discord
    dg.AddHandler(messageHandler) // Certifique-se de que isso está aqui
}

// Função para lidar com mensagens
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    fmt.Println("Mensagem recebida:", m.Content) // Log da mensagem recebida
    if m.Author.ID == s.State.User.ID {
        return
    }
    if m.Content == "!pr" {
        fmt.Println("Comando !pr recebido") // Log do comando
        // Criação do embed com botões
        embed := &discordgo.MessageEmbed{
            Title:       "Confirmação",
            Description: "Você deseja continuar?",
            Color:       0x00ff00,
        }
        // Envio do embed com botões
        msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
        if err != nil {
            fmt.Println("Error sending embed: ", err)
            return
        }
        // Adicionando reações ao embed
        _ = s.MessageReactionAdd(m.ChannelID, msg.ID, "✅")
        _ = s.MessageReactionAdd(m.ChannelID, msg.ID, "❌")
    }
}
