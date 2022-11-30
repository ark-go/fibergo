package tbot

import (
	"log"
	"os"
	"time"

	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

func (b *Bot) sendPhoto(usr *userdata.User, filePath string, delay ...int) (*telegrambot.Message, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	msgout, err := b.Api.SendPhoto(&telegrambot.SendPhotoParams{
		ChatID: usr.ChatId,
		Photo: &telegrambot.FileReader{
			Name:   "test",
			Reader: file,
		},
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	delayTime := 15
	if len(delay) > 0 && delay[0] > 5 {
		delayTime = delay[0]
	}
	usr.LastMsgQueues.Add(int64(msgout.MessageID), int64(usr.ChatId), delayTime)

	if err == nil && len(delay) > 0 {
		go func() {
			time.Sleep(time.Duration(delayTime) * time.Second)
			//b.deleteMessage(msgout)
			b.deleteMessage(usr)
		}()
	}
	return msgout, nil
}
