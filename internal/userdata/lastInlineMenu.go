package userdata

import (
	"time"

	"github.com/nickname76/telegrambot"
)

type LastInlineMenu struct {
	// последнее сообщение-меню отправленное клиенту
	MessageID  *telegrambot.MessageID
	ChatID     *telegrambot.ChatID
	FileId     string // фото
	TimeDelete time.Time
}

type LastInlineMenuAll map[telegrambot.ChatID]*LastInlineMenu

func (l LastInlineMenuAll) Add(chatId telegrambot.ChatID, messageID telegrambot.MessageID, fileId string, delayTime int) {
	l[chatId] = &LastInlineMenu{
		MessageID:  &messageID,
		ChatID:     &chatId,
		FileId:     fileId,
		TimeDelete: time.Now().Add(time.Duration(delayTime) * time.Second),
	}
}

func (l LastInlineMenuAll) Drop(chatId telegrambot.ChatID) {
	delete(l, chatId)
}

func (l LastInlineMenuAll) GetFileId(chatId telegrambot.ChatID) (fileId string) {
	if v, ok := l[chatId]; ok {
		return v.FileId
	}
	return ""
}
