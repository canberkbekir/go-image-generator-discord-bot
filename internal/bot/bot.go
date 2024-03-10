package bot

import (
	"github.com/bwmarrin/discordgo"
	"go-image-generator-discord-bot/internal/chatgpt"
	"go-image-generator-discord-bot/internal/commands"
	util "go-image-generator-discord-bot/internal/utils"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type DiscordBotRunner interface {
	Run() error
}

type Bot struct {
}

var KeyBot string

func (b Bot) Run() error {
	dg, err := discordgo.New("Bot " + KeyBot)
	if err != nil {
		util.Logger.Error().Err(err)
		return err
	}

	//Handlers
	dg.AddHandler(messageCreateHandler)

	err = dg.Open()
	if err != nil {
		util.Logger.Error().Err(err)
		return err
	}
	// Wait here until CTRL-C or other term signal is received.
	util.Logger.Info().Msg("Kid CB is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		util.Logger.Error().Err(err)
		return err
	}

	return nil
}

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	ai := chatgpt.NewOpenAI()
	imageCommands := commands.NewImageCommand(ai)

	args := strings.Split(m.Content, " ")

	switch args[0] {
	case "!image":
		err := imageCommands.CreateImage(s, m)
		if err != nil {
			return
		}
	}
}

func NewDiscordBotRunner() DiscordBotRunner {
	return &Bot{}
}
