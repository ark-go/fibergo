package tbot

import (
	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

func (b *Bot) sendLocation(user *userdata.User) {
	b.Api.SendLocation(&telegrambot.SendLocationParams{
		ChatID:             user.MsgChatId(),
		Latitude:           59.8983,
		Longitude:          30.2618,
		HorizontalAccuracy: 10.4, // радиус в метрах
		LivePeriod:         60,   // В секундах от 60
	})
}
