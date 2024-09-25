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

    // Adiciona um handler para o evento de mensagem
    dg.AddHandler(messageCreate)

    // Conecta o bot ao Discord
    err = dg.Open()
    if err != nil {
        fmt.Println("Error opening connection: ", err)
        return
    }

    fmt.Println("Bot is now running. Press CTRL+C to exit.")
    select {}
}

// Função para lidar com o evento de mensagem
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

    // Verifica se a mensagem começa com o prefixo
    if m.Content == prefix + "ping" {
        // Envia a mensagem de resposta
        _, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
    }
}
