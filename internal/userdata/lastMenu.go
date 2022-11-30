package userdata

import (
	"time"

	"github.com/nickname76/telegrambot"
)

type LastMenu struct {
	// последнее сообщение-меню отправленное клиенту
	MessageID  telegrambot.MessageID
	ChatID     telegrambot.ChatID
	TimeDelete time.Time
}

type LastMenuAll map[telegrambot.ChatID]LastMenu

func (l *LastMenuAll) Add(chatId telegrambot.ChatID, messageID telegrambot.MessageID, delayTime int) {
	(*l)[chatId] = LastMenu{
		MessageID:  messageID,
		ChatID:     chatId,
		TimeDelete: time.Now().Add(time.Duration(delayTime) * time.Second),
	}
}
