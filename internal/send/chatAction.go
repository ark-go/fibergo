package send

import "github.com/nickname76/telegrambot"

func (s *Send) ChatAction() {
	s.api.SendChatAction(&telegrambot.SendChatActionParams{
		ChatID: s.User.GetChatId(),
		Action: telegrambot.ChatActionUploadPhoto,
	})
}
