package pDefault

import "github.com/nickname76/telegrambot"

func (b *ProgDef) MenuButton() [][]*telegrambot.KeyboardButton {

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
			Text:            "Нажми сюда",
			RequestLocation: true,
		},
		{

			Text: "Сброс",
		},
	}}
}
