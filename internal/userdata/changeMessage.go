package userdata

import "github.com/nickname76/telegrambot"

func (u *User) ChangeMessage(msg *telegrambot.Message) {

	//if u.ClientType == Client_Channel {

	// u.UserId = msg.From.ID
	// u.ChatId = msg.Chat.ID
	//	u.Msg = msg
	u.checkMessageType(msg)
	u.checkClientType(msg)
}
