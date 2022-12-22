package send

import "github.com/nickname76/telegrambot"

func (s *Send) AnswerCallbackQuery(text string, alert ...bool) {
	s.api.AnswerCallbackQuery(&telegrambot.AnswerCallbackQueryParams{
		CallbackQueryID: s.User.Info.CallbackQueryID,
		Text:            text,
		ShowAlert:       len(alert) > 0, // выводить в алерт
		CacheTime:       2,              // кеширует и не запрашивает сервер
	})
}
