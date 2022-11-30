package tbot

import (
	"log"

	"github.com/nickname76/telegrambot"
)

func (b *Bot) editInlineMessage(inl *telegrambot.CallbackQuery, message string) error {
	log.Println("inl id", inl.InlineMessageID, " :>", inl.Message.MessageID)
	b.Api.EditMessageText(&telegrambot.EditMessageTextParams{
		ChatID:          inl.Message.Chat.ID,
		MessageID:       inl.Message.MessageID,
		InlineMessageID: inl.InlineMessageID,
		Text:            message,
		ReplyMarkup:     inl.Message.ReplyMarkup,
	})
	return nil
}

// editMessageReplyMarkup  вроде только кнопки
