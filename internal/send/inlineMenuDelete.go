package send

import (
	"log"

	"github.com/nickname76/telegrambot"
)

// удаляем  InlineMenu если оно было
func (s *Send) InlineMenuDelete() error {
	// log.Printf("all inl %++v\n", s.User.UserData.InlineMenuAll)
	for _, inlineMenu := range s.User.UserData.InlineMenuAll {
		if *inlineMenu.ChatID == s.User.GetChatId() {
			if inlineMenu.MessageID != nil {
				log.Println("Удаляем наше сообщение с кнопками", *inlineMenu.ChatID, *inlineMenu.MessageID)
				return s.api.DeleteMessage(&telegrambot.DeleteMessageParams{
					ChatID:    *inlineMenu.ChatID,
					MessageID: *inlineMenu.MessageID,
				})
			}
		}
	}
	return nil
}
