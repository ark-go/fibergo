package programs

import (
	"strings"

	"github.com/nickname76/telegrambot"
)

/*
была ли команда со слешем /cmd , (которую мы ждем)

	извлекает из сообщения команду, если она там есть и сравнивает с параметром
	результат подтверждение bool
*/
func (p *Program) IsCommand(cmd string) bool {
	command, _ := p.ParseMessageCommand()
	return command == cmd
}

/*
парсер команд посланых боту  command@arguments

	возвращает команду и аргументы раздельно
*/
func (p *Program) ParseMessageCommand() (command string, args string) {
	var (
		text         string
		textEntities []*telegrambot.MessageEntity
	)

	switch {
	case p.Send.User.Msg.Text != "":
		text = p.Send.User.Msg.Text
		textEntities = p.Send.User.Msg.Entities
	case p.Send.User.Msg.Caption != "":
		text = p.Send.User.Msg.Caption
		textEntities = p.Send.User.Msg.CaptionEntities
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
