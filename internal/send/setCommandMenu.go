package send

import "github.com/nickname76/telegrambot"

// установка текстового меню у пользователя или в группе в строке редактирования
//
//	если параметр nil - установится одна кнопка /start
func (s *Send) SetCommandMenu(botCommand []*telegrambot.BotCommand) {
	if botCommand == nil {
		botCommand = []*telegrambot.BotCommand{
			{
				Command:     "/start",
				Description: "Старт или сброс",
			},
		}
	}
	s.api.SetMyCommands(&telegrambot.SetMyCommandsParams{
		Commands: botCommand,
		Scope: &telegrambot.BotCommandScope{
			//	Type:   telegrambot.BotCommandScopeTypeAllChatAdministrators, // в чаты где бот адмистратором
			Type:   telegrambot.BotCommandScopeTypeChat, // в чат откуда вылез пользователь
			ChatID: s.User.GetChatId(),
			UserID: s.User.GetUserId(),
		},
	})
}
