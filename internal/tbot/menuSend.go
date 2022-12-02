package tbot

import (
	"log"
	"time"

	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

func (b *Bot) MenuSend(user *userdata.User, text string) {

	msgout, err := b.Api.SendMessage(&telegrambot.SendMessageParams{
		ChatID:                user.MsgChatId(),
		Text:                  text,
		ProtectContent:        true, // Запрет копирования - пересылки сообщения
		DisableWebPagePreview: true, // отключить предпросмотр ссылок url
		DisableNotification:   true, // без звука
		ReplyMarkup: &telegrambot.ReplyKeyboardMarkup{
			OneTimeKeyboard: true,           // возможность сворачивать
			ResizeKeyboard:  true,           // клавиатура меньше по вертикали
			Keyboard:        b.MenuButton(), // кнопки
		},
	})
	if err != nil {
		log.Println("Не удалось отправить меню кнопки")
	} else {
		user.Last.MenuAll.Add(msgout.Chat.ID, msgout.MessageID, 15)
		log.Printf("Отправили меню %+v", user.Last.MenuAll)
		go func() {
			time.Sleep(time.Duration(16) * time.Second)
			log.Printf("Хотим удалить меню %+v", user.Last.MenuAll)
			b.deleteMessageMenu(user)
		}()
	}
}
