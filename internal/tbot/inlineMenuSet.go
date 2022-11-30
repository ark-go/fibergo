package tbot

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/ark-go/fibergo/internal/utils"
	"github.com/nickname76/telegrambot"
)

// Установка InlineMenu меню к сообщению
func (b *Bot) InlineMenuSet(user *userdata.User, knopki *telegrambot.InlineKeyboardMarkup) {
	path, err1 := utils.GetExecPath()
	if err1 != nil {
		log.Println(err1)
		return
	}
	b.InlineMenuDelete(user)
	kto := 2
	var inlMsg *telegrambot.Message
	var err error
	if kto == 1 {
		inlMsg, err = b.inlineMenuText(user, knopki)
	} else if kto == 2 {

		path = filepath.Join(path, "img-210.png")
		inlMsg, err = b.inlineMenuPhoto(user, knopki, path)
	} else {
		path = filepath.Join(path, "gif-1.gif")
		inlMsg, err = b.inlineMenuAnimate(user, knopki, path)
	}

	if err != nil {
		log.Println("Ошибка отправки меню", err.Error())
		return
	}
	user.Last.InlineMenuAll.Add(user.MsgChatId(), inlMsg.MessageID, 0)
	log.Println("Отправили меню")
}

// меню только с текстом
func (b *Bot) inlineMenuText(user *userdata.User, knopki *telegrambot.InlineKeyboardMarkup) (*telegrambot.Message, error) {
	msg, err := b.Api.SendMessage(&telegrambot.SendMessageParams{
		ChatID: user.MsgChatId(),
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
func (b *Bot) inlineMenuPhoto(user *userdata.User, knopki *telegrambot.InlineKeyboardMarkup, filePath string) (*telegrambot.Message, error) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	msg, err := b.Api.SendPhoto(&telegrambot.SendPhotoParams{
		ChatID: user.MsgChatId(),
		Photo: &telegrambot.FileReader{
			Name:   "test",
			Reader: file,
		},
		ProtectContent: true,
		ReplyMarkup:    knopki,
		ParseMode:      telegrambot.ParseModeHTML,
		Caption:        "<b>Я тут бот.</b>\n" + user.Txt,
	})
	return msg, err
}

// меню с анимацией
func (b *Bot) inlineMenuAnimate(user *userdata.User, knopki *telegrambot.InlineKeyboardMarkup, filePath string) (*telegrambot.Message, error) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	msg, err := b.Api.SendAnimation(&telegrambot.SendAnimationParams{
		ChatID: user.MsgChatId(),
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
