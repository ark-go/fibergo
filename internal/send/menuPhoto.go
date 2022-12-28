package send

import (
	"log"
	"path/filepath"

	"github.com/ark-go/fibergo/internal/utils"
	"github.com/nickname76/telegrambot"
)

func (s *Send) MenuPhoto(knopki *telegrambot.InlineKeyboardMarkup, filePhoto, fileId string, caption ...string) {
	s.InlineMenuDelete()

	path := filepath.Join(utils.ExecDir, filePhoto)

	inlMsg, err := s.Photo(knopki, path, fileId, caption...)
	if err != nil {
		log.Println("Ошибка отправки меню", err.Error())
		return
	}

	s.User.UserData.InlineMenuAll.Add(s.User.GetChatId(), inlMsg.MessageID, string(inlMsg.Photo[0].FileID), 0)
}
