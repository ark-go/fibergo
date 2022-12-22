package send

import (
	"time"

	// "github.com/ark-go/fibergo/internal/tbot"  "send" cycle ??
	"github.com/nickname76/telegrambot"
)

// удаляет сообщение msg и отсылает свое message
// ставит отправленные сообщения в очередь
// delay - можно задать задержку в сек. когда считать сообщение устаревшим для удаления
// не меньше 5 сек..
func (s *Send) SendTimeMessage(message string, delay ...int) {
	delayTime := 10
	if len(delay) > 0 && delay[0] >= 5 {
		delayTime = delay[0]
	}
	s.DeleteMessage()
	if message != "" {
		msgout, err := s.api.SendMessage(&telegrambot.SendMessageParams{
			ChatID:                s.User.GetChatId(),
			Text:                  message,
			ParseMode:             telegrambot.ParseModeHTML,
			ProtectContent:        true, // Запрет копирования - пересылки сообщения
			DisableWebPagePreview: true, // отключить предпросмотр ссылок url
			DisableNotification:   true, // без звука
		})
		s.User.UserData.MsgQueues.Add(int64(msgout.MessageID), int64(s.User.GetChatId()), delayTime)
		// ! спорный вопрос, пока для теста, запуск счетчика для удаления
		// вероятно он будет есть память не отпуская " s "
		if err == nil {
			go func() {
				time.Sleep(time.Duration(delayTime) * time.Second)

				s.DeleteMessage()
			}()
		}
	}

}
