package send

import (
	"strings"

	"github.com/ark-go/fibergo/internal/msgtypes"
	"github.com/nickname76/telegrambot"
)

// Замена Caption у MessagePhoto
// заменяем текст у картинки
func (s *Send) MenuCaption(replyMark *telegrambot.InlineKeyboardMarkup, mess string) (*telegrambot.Message, error) {

	if msg, ok := s.User.UserData.InlineMenuAll[s.User.GetChatId()]; ok {
		mess, err := s.api.EditMessageCaption(&telegrambot.EditMessageCaptionParams{
			ChatID:      s.User.GetChatId(),
			MessageID:   *msg.MessageID,
			Caption:     mess,
			ParseMode:   telegrambot.ParseModeHTML,
			ReplyMarkup: replyMark,
		})
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return nil, msgtypes.ErrNotFound
			}
			if strings.Contains(err.Error(), "not modified") {
				return nil, msgtypes.ErrNotModified
			}
			return nil, err
		}
		return mess, nil
	}

	return nil, msgtypes.ErrNotFoundKeyUser
}
