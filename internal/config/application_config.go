package config

import (
	"go-image-generator-discord-bot/internal/bot"
	"go-image-generator-discord-bot/internal/chatgpt"
	"os"
)

func InitConfiguration() {
	bot.KeyBot = os.Getenv("DISCORD_KEY")
	chatgpt.KeyOpenAI = os.Getenv("OPENAI_KEY")
}
