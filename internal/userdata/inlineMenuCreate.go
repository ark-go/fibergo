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
				Text:         "333333",
				CallbackData: "333333",
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
				Text:         "Сброс",
				CallbackData: "Сброс",
			},
		}},
	}
}
