package tbot

import (
	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

// удаляем  InlineMenu если оно было
func (b *Bot) InlineMenuDelete(user *userdata.User) error {
	for _, inlineMenu := range user.LastInlineMenuAll {
		if *inlineMenu.ChatID == user.ChatId {
			if inlineMenu.MessageID != nil {
				return b.Api.DeleteMessage(&telegrambot.DeleteMessageParams{
					ChatID:    *inlineMenu.ChatID,
					MessageID: *inlineMenu.MessageID,
				})
			}
		}
	}
	return nil
}