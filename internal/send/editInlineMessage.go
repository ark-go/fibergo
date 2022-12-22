package send

import (
	"log"

	"github.com/nickname76/telegrambot"
)

func (s *Send) editInlineMessage(inl *telegrambot.CallbackQuery, message string) error {
	log.Println("inl id", inl.InlineMessageID, " :>", inl.Message.MessageID)
	s.api.EditMessageText(&telegrambot.EditMessageTextParams{
		ChatID:          inl.Message.Chat.ID,
		MessageID:       inl.Message.MessageID,
		InlineMessageID: inl.InlineMessageID,
		Text:            message,
		ReplyMarkup:     inl.Message.ReplyMarkup,
	})
	return nil
}

// editMessageReplyMarkup  вроде только кнопки
