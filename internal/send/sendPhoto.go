package send

import (
	"log"
	"os"
	"time"

	"github.com/nickname76/telegrambot"
)

func (s *Send) sendPhoto(filePath string, delay ...int) (*telegrambot.Message, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	msgout, err := s.api.SendPhoto(&telegrambot.SendPhotoParams{
		ChatID: s.User.GetChatId(),
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
	s.User.UserData.MsgQueues.Add(int64(msgout.MessageID), int64(s.User.GetChatId()), delayTime)

	if err == nil && len(delay) > 0 {
		go func() {
			time.Sleep(time.Duration(delayTime) * time.Second)
			s.DeleteMessage()
		}()
	}
	return msgout, nil
}
