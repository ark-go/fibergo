package tbot

import "github.com/nickname76/telegrambot"

func (b *Bot) MenuButton() [][]*telegrambot.KeyboardButton {

	return [][]*telegrambot.KeyboardButton{{
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
			Text:            "Нажми сюда",
			RequestLocation: true,
		},
		{

			Text: "Сброс",
		},
	}}
}
