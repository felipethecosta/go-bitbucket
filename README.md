# Documentação Golang Integration Discord and Bitbucket

## Informações do Projeto

**Nome do Projeto:** go-bitbucket
**Descrição:** Integração do Discord com o Bitbucket usando um bot em Go.

## Arquivos do Projeto

1. **main.go**
   - Contém a lógica principal do bot, incluindo a inicialização da sessão do Discord, envio de mensagens e tratamento de eventos de mensagens.

2. **.gitignore**
   - Ignora o arquivo de variáveis de ambiente `.env` para evitar que informações sensíveis sejam versionadas.

3. **go.mod**
   - Define o módulo do Go e as dependências necessárias para o projeto, incluindo `discordgo` e `godotenv`.

4. **go.sum**
   - Contém as somas de verificação das dependências para garantir a integridade do código.

## Estrutura do Projeto

go-bitbucket/
├── .gitignore
├── go.mod
├── go.sum
└── main.go


## Dependências

- `github.com/bwmarrin/discordgo`: Biblioteca para interagir com a API do Discord.
- `github.com/joho/godotenv`: Carrega variáveis de ambiente de um arquivo `.env`.

## Instruções de Uso

1. **Configuração do Ambiente:**
   - Crie um arquivo `.env` na raiz do projeto com as variáveis `DISCORD_TOKEN` e `CHANNEL_ID`.

2. **Execução do Bot:**
   - Execute o comando `go run main.go` para iniciar o bot.

3. **Comandos do Bot:**
   - O bot responde ao comando `!ping` com `Pong!`.

## Observações

- Certifique-se de que as dependências estão instaladas executando `go mod tidy`.
- O bot deve ter permissões adequadas no canal do Discord para enviar mensagens.
