package tbot

import (
	"log"

	"github.com/nickname76/telegrambot"
)

func (b *Bot) inlineCommand(inl *telegrambot.CallbackQuery) error {
	log.Println("Команда Inline кнопки:", inl.Data) // , inl.Message.Text

	// Отправка информацции в ответ о принятой команде
	b.Api.AnswerCallbackQuery(&telegrambot.AnswerCallbackQueryParams{
		CallbackQueryID: inl.ID,
		Text:            inl.Message.Text + " : " + inl.Data,
		//ShowAlert: true, // выводить в алерт
		CacheTime: 10, // кеширует и не запрашивает сервер
	})
	b.editInlineMessage(inl, "🌲Нажал : "+inl.Data)
	return nil
}
