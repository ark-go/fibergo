package send

import (
	"log"
	// "time"

	"github.com/nickname76/telegrambot"
)

func (s *Send) InlineCommand(inl *telegrambot.CallbackQuery) error {
	log.Println("Команда Inline кнопки:", inl.Data, inl.Message.Text, inl.ChatInstance) // , inl.Message.Text
	// time.Sleep(25 * time.Second) // без ответа AnswerCallbackQuery  бот больше не будет получать собщения
	// Отправка информацции в ответ о принятой команде
	s.api.AnswerCallbackQuery(&telegrambot.AnswerCallbackQueryParams{
		CallbackQueryID: inl.ID,
		Text:            inl.Message.Text + " : " + inl.Data,
		//ShowAlert: true, // выводить в алерт
		CacheTime: 10, // кеширует и не запрашивает сервер
	})
	s.editInlineMessage(inl, "🌲Нажал : "+inl.Data)
	return nil
}
