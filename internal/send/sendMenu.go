package send

import (
	"log"

	"github.com/nickname76/telegrambot"
)

func (s *Send) SendMenu(button [][]*telegrambot.KeyboardButton, text string) {
	s.deleteMessageMenu()
	msgout, err := s.api.SendMessage(&telegrambot.SendMessageParams{
		ChatID:                s.User.GetChatId(),
		Text:                  text,
		ProtectContent:        true, // Запрет копирования - пересылки сообщения
		DisableWebPagePreview: true, // отключить предпросмотр ссылок url
		DisableNotification:   true, // без звука
		ReplyMarkup: &telegrambot.ReplyKeyboardMarkup{
			OneTimeKeyboard: true,   // возможность сворачивать
			ResizeKeyboard:  true,   // клавиатура меньше по вертикали
			Keyboard:        button, //b.MenuButton(), // кнопки
		},
	})
	if err != nil {
		log.Println("Не удалось отправить меню кнопки")
	} else {
		s.User.UserData.MenuAll.Add(msgout.Chat.ID, msgout.MessageID, 15)
		// go func() {
		// 	time.Sleep(time.Duration(16) * time.Second)
		// 	s.deleteMessageMenu()
		// }()
	}
}
