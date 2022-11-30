package userdata

import "github.com/nickname76/telegrambot"

func (u *User) ChangeMessage(msg *telegrambot.Message) {
	u.TgUser = msg.From
	u.UserId = msg.From.ID
	u.ChatId = msg.Chat.ID
	//	u.Msg = msg
	u.checkTypeMessage(msg)
	u.checkTypeClient(msg)
}
