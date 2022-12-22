package userdata

import (
	"time"

	"github.com/nickname76/telegrambot"
)

type LastInlineMenu struct {
	// последнее сообщение-меню отправленное клиенту
	MessageID  *telegrambot.MessageID
	ChatID     *telegrambot.ChatID
	TimeDelete time.Time
}

type LastInlineMenuAll map[telegrambot.ChatID]*LastInlineMenu

func (l LastInlineMenuAll) Add(chatId telegrambot.ChatID, messageID telegrambot.MessageID, delayTime int) {
	l[chatId] = &LastInlineMenu{
		MessageID:  &messageID,
		ChatID:     &chatId,
		TimeDelete: time.Now().Add(time.Duration(delayTime) * time.Second),
	}
}

func (l LastInlineMenuAll) Drop(chatId telegrambot.ChatID) {
	delete(l, chatId)
}
