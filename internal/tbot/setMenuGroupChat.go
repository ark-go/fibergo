package tbot

import (
	"log"
	"time"

	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

func (b *Bot) setMenuGroupChat(user *userdata.User) {

	msgout, err := b.Api.SendMessage(&telegrambot.SendMessageParams{
		ChatID: user.ChatId,
		//Text:                  fmt.Sprintf("Привет %v,  %v, ваш ид: %d", msg.From.FirstName, msg.From.LastName, msg.Chat.ID),
		Text:                  ".",
		ProtectContent:        true, // Запрет копирования - пересылки сообщения
		DisableWebPagePreview: true, // отключить предпросмотр ссылок url
		DisableNotification:   true, // без звука
		ReplyMarkup: &telegrambot.ReplyKeyboardMarkup{
			OneTimeKeyboard: true, // возможность сворачивать
			ResizeKeyboard:  true, // клавиатура меньше по вертикали
			Keyboard: [][]*telegrambot.KeyboardButton{{
				{
					Text: "Hello",
				},

				{
					// WebApp: &telegrambot.WebAppInfo{
					// 	URL: "https://bake.x.arkadii.ru",
					// },
					Text: "Лах",
				},
			}, {
				{
					// WebApp: &telegrambot.WebAppInfo{
					// 	URL: "https://bake.x.arkadii.ru",
					// },
					Text: "Лах",
				},
				{

					Text: "Сброс",
				},
			}},
		},
	})
	if err != nil {
		log.Println("Не удалось отправить меню кнопки")
	} else {
		//user.LastMenu.MessageID = msgout.MessageID
		//user.LastMenu.ChatID = msgout.Chat.ID
		user.LastMenuAll.Add(msgout.Chat.ID, msgout.MessageID, 15)
		log.Printf("Отправили меню %+v", user.LastMenuAll)
		go func() {
			time.Sleep(time.Duration(16) * time.Second)
			log.Printf("Хотим удалить меню %+v", user.LastMenuAll)
			b.deleteMessageMenu(user)
		}()
	}
}
