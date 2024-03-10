package commands

import (
	"bytes"
	"github.com/bwmarrin/discordgo"
	"go-image-generator-discord-bot/internal/chatgpt"
	"strings"
)

type ImageCommandInterface interface {
	CreateImage(s *discordgo.Session, m *discordgo.MessageCreate) error
}

type ImageCommand struct {
	ai chatgpt.OpenAIRunner
}

func (ic ImageCommand) CreateImage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	args := strings.SplitN(m.Content, " ", 2)
	prompt := args[1]

	imgBytes, err := ic.ai.CreateImage(prompt)
	if err != nil {
		return err
	}

	r := bytes.NewReader(*imgBytes)

	_, err = s.ChannelFileSend(m.ChannelID, "image.png", r)
	return nil
}

func CreateImage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

}

func NewImageCommand(ai chatgpt.OpenAIRunner) ImageCommandInterface {
	return &ImageCommand{
		ai: ai,
	}
}
