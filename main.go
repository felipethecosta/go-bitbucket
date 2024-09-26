package main

import (
	"log"
	"os"

	"go-bitbucket/services/discord/commands"
	"go-bitbucket/services/discord/handlers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

   var Token string
   var ChannelID string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	Token = os.Getenv("DISCORD")
	if Token == "" {
		log.Fatal("Token não encontrado. Verifique se a variável de ambiente DISCORD está definida.")
	}

	ChannelID = os.Getenv("CHANNEL_ID")
	if ChannelID == "" {
		log.Fatal("Channel ID não encontrado. Verifique se a variável de ambiente CHANNEL_ID está definida.")
	}

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("Erro ao criar a sessão: %v", err)
	}

	// Adiciona um handler para mensagens
	dg.AddHandler(commands.PingCommand)
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão: %v", err)
	}

	log.Println("Bot está online!")
	select {}
}

// messageCreate é chamado sempre que uma nova mensagem é criada
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("Mensagem recebida: %s de %s no canal %s", m.Content, m.Author.Username, m.ChannelID)

	if m.Content == "!embed" {
		handlers.SendEmbedMessage(s, ChannelID)
		log.Println("Enviando mensagem embed.")
	}
	if m.Content == "!react" {
		handlers.AddReaction(s, m.ChannelID, m.ID, "👍")
		log.Println("Adicionando reação.")
	}
	if m.Content == "!ping" {
		commands.PingCommand(s, m)
		log.Println("Respondendo ao comando ping.")
	}
	if m.Content == "!help" {
		commands.HelpCommand(s, m)
		log.Println("Respondendo ao comando help.")
	}
}
