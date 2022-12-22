package send

import (
	"log"

	"github.com/nickname76/telegrambot"
)

// Удаляем все сообщение отправленные пользователем в текущий чат
// указывая кол-во например  10 последних сообщений
func (s *Send) DeleteLastMessages(count int64) {
	var msgId int64 = int64(s.User.Msg.MessageID)
	for i := msgId; i > msgId-count; i-- {
		s.api.DeleteMessage(&telegrambot.DeleteMessageParams{
			ChatID:    s.User.Msg.Chat.ID,
			MessageID: telegrambot.MessageID(i),
		})
	}
	log.Println("На удаление последних", count, "сообщений.")
}
