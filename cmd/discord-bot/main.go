package main

import (
	"go-image-generator-discord-bot/internal/bot"
	"go-image-generator-discord-bot/internal/config"
	util "go-image-generator-discord-bot/internal/utils"
)

const prefix string = "!"

func main() {
	config.InitConfiguration()
	util.InitLog()

	err := bot.NewDiscordBotRunner().Run()
	if err != nil {
		util.Logger.Error().Err(err)
	}
}
