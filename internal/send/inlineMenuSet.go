package send

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/ark-go/fibergo/internal/utils"
	"github.com/nickname76/telegrambot"
)

// Установка InlineMenu меню к сообщению
func (s *Send) InlineMenuSet(knopki *telegrambot.InlineKeyboardMarkup, text ...string) {
	path, err1 := utils.GetExecPath()
	if err1 != nil {
		log.Println(err1)
		return
	}
	s.InlineMenuDelete()
	kto := 2
	var inlMsg *telegrambot.Message
	var err error
	if kto == 1 {
		inlMsg, err = s.inlineMenuText(s.User, knopki)
	} else if kto == 2 {

		path = filepath.Join(path, "img-210.png")
		inlMsg, err = s.inlineMenuPhoto(knopki, path, text...)
	} else {
		path = filepath.Join(path, "gif-1.gif")
		inlMsg, err = s.inlineMenuAnimate(s.User, knopki, path)
	}

	if err != nil {
		log.Println("Ошибка отправки меню", err.Error())
		return
	}
	log.Println("inl:", inlMsg.MessageID, s.User.GetChatId(), inlMsg.Chat.ID)
	s.User.UserData.InlineMenuAll.Add(s.User.GetChatId(), inlMsg.MessageID, 0)
	log.Println("inl2:", *s.User.UserData.InlineMenuAll[s.User.GetChatId()].MessageID)
}

// меню только с текстом
func (s *Send) inlineMenuText(user *userdata.User, knopki *telegrambot.InlineKeyboardMarkup) (*telegrambot.Message, error) {
	msg, err := s.api.SendMessage(&telegrambot.SendMessageParams{
		ChatID: user.GetChatId(),
		//Text:                  fmt.Sprintf("Привет %v,  %v, ваш ид: %d", msg.From.FirstName, msg.From.LastName, msg.Chat.ID),
		Text:                  "Меню:",
		ProtectContent:        true, // Запрет копирования - пересылки сообщения
		DisableWebPagePreview: true, // отключить предпросмотр ссылок url
		DisableNotification:   true, // без звука
		ReplyMarkup:           knopki,
	})
	return msg, err
}

// меню с фоткой
func (s *Send) inlineMenuPhoto(knopki *telegrambot.InlineKeyboardMarkup, filePath string, caption ...string) (*telegrambot.Message, error) {
	if len(caption) == 0 {
		caption = append(caption, "<b>Я тут бот.</b>\n")
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	msg, err := s.api.SendPhoto(&telegrambot.SendPhotoParams{
		ChatID: s.User.GetChatId(),
		Photo: &telegrambot.FileReader{
			Name:   "test",
			Reader: file,
		},
		ProtectContent: true,
		ReplyMarkup:    knopki,
		ParseMode:      telegrambot.ParseModeHTML,
		Caption:        caption[0],
	})

	if err == nil {
		id := msg.Photo[0].FileID // этот ИД можно отправить вместо картинки?
		log.Println("Картинка отправлена ID:", id)
	}
	return msg, err
}

// меню с анимацией
func (s *Send) inlineMenuAnimate(user *userdata.User, knopki *telegrambot.InlineKeyboardMarkup, filePath string) (*telegrambot.Message, error) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	msg, err := s.api.SendAnimation(&telegrambot.SendAnimationParams{
		ChatID: user.GetChatId(),
		Animation: &telegrambot.FileReader{
			Name:   "test.gif",
			Reader: file,
		},
		//Duration:       100,
		Width:          800,
		ProtectContent: true,
		ReplyMarkup:    knopki,
		ParseMode:      telegrambot.ParseModeHTML,
		Caption:        "<b>А я тут бот.</b>\n" + user.Txt,
	})
	return msg, err
}
