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
	b.InlineMenuDelete(user)
	kto := 11
	var inlMsg *telegrambot.Message
	var err error
	if kto == 1 {
		inlMsg, err = b.inlineMenuText(user, knopki)
	} else {
		path, err1 := utils.GetExecPath()
		if err1 != nil {
			log.Println(err1)
			return
		}
		path = filepath.Join(path, "img-1.jpg")
		inlMsg, err = b.inlineMenuPhoto(user, knopki, path)
	}

	if err != nil {
		log.Println("Ошибка отправки меню", err.Error())
		return
	}
	user.LastInlineMenuAll.Add(user.ChatId, inlMsg.MessageID, 0)
	log.Println("Отправили меню")
}

func (b *Bot) inlineMenuText(user *userdata.User, knopki *telegrambot.InlineKeyboardMarkup) (*telegrambot.Message, error) {
	msg, err := b.Api.SendMessage(&telegrambot.SendMessageParams{
		ChatID: user.ChatId,
		//Text:                  fmt.Sprintf("Привет %v,  %v, ваш ид: %d", msg.From.FirstName, msg.From.LastName, msg.Chat.ID),
		Text:                  "Меню:",
		ProtectContent:        true, // Запрет копирования - пересылки сообщения
		DisableWebPagePreview: true, // отключить предпросмотр ссылок url
		DisableNotification:   true, // без звука
		ReplyMarkup:           knopki,
	})
	return msg, err
}

func (b *Bot) inlineMenuPhoto(user *userdata.User, knopki *telegrambot.InlineKeyboardMarkup, filePath string) (*telegrambot.Message, error) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	msg, err := b.Api.SendPhoto(&telegrambot.SendPhotoParams{
		ChatID: user.ChatId,
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
