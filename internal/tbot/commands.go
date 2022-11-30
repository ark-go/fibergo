package tbot

import (
	"strings"

	"github.com/nickname76/telegrambot"
)

func (b *Bot) ParseMessageCommand(msg *telegrambot.Message) (command string, args string) {
	var (
		text         string
		textEntities []*telegrambot.MessageEntity
	)

	switch {
	case msg.Text != "":
		text = msg.Text
		textEntities = msg.Entities
	case msg.Caption != "":
		text = msg.Caption
		textEntities = msg.CaptionEntities
	default:
		return
	}

	for _, entity := range textEntities {
		if entity.Type != telegrambot.MessageEntityTypeBotCommand || entity.Offset != 0 {
			continue
		}

		command = text[1:entity.Length]

		usernameIndex := strings.Index(command, "@")
		if usernameIndex != -1 {
			command = command[:usernameIndex]
		}

		args = strings.TrimSpace(text[entity.Length:])

		break
	}

	return
}
