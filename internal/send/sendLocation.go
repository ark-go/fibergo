package send

import (
	"github.com/nickname76/telegrambot"
)

func (s *Send) sendLocation() {
	s.api.SendLocation(&telegrambot.SendLocationParams{
		ChatID:             s.User.GetChatId(),
		Latitude:           59.8983,
		Longitude:          30.2618,
		HorizontalAccuracy: 10.4, // радиус в метрах
		LivePeriod:         60,   // В секундах от 60
	})
}
