package userdata

import "github.com/nickname76/telegrambot"

func (u *User) CreateInlineMenu() *telegrambot.InlineKeyboardMarkup {
	return &telegrambot.InlineKeyboardMarkup{
		InlineKeyboard: [][]*telegrambot.InlineKeyboardButton{{
			{
				Text:         "Привет",
				CallbackData: "```Привет```",
			},
			{
				Text: "www",
				WebApp: &telegrambot.WebAppInfo{
					URL: "https://bake.x.arkadii.ru",
				},
			},
			{
				Text:         "1",
				CallbackData: "1",
			},
			{
				Text:         "Привет",
				CallbackData: "32",
			},
		}, {
			{
				Text:         "Тут",
				CallbackData: "Я вася@петров",
			},
			{
				Text:         "23",
				CallbackData: "23",
			},
			{
				Text:         "233",
				CallbackData: "2334",
			},
			{
				Text:         "233",
				CallbackData: "2334",
			},
			{
				Text:         "Сброс",
				CallbackData: "Сброс",
			},
		}},
	}
}
