package tbot

import (
	"time"

	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

// удаляет сообщение msg и отсылает свое message
func (b *Bot) sendTimeMessage(usr *userdata.User, message string, delay ...int) {
	delayTime := 10
	if len(delay) > 0 && delay[0] > 5 {
		delayTime = delay[0]
	}
	b.deleteMessage(usr)
	if message != "" {
		msgout, err := b.Api.SendMessage(&telegrambot.SendMessageParams{
			ChatID:                usr.ChatId,
			Text:                  message,
			ParseMode:             telegrambot.ParseModeHTML,
			ProtectContent:        true, // Запрет копирования - пересылки сообщения
			DisableWebPagePreview: true, // отключить предпросмотр ссылок url
			DisableNotification:   true, // без звука
		})
		usr.LastMsgQueues.Add(int64(msgout.MessageID), int64(usr.ChatId), delayTime)
		if err == nil {
			go func() {
				time.Sleep(time.Duration(delayTime) * time.Second)
				b.deleteMessage(usr)
			}()
		}
	}

}
