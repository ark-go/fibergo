package stages

import (
	"log"

	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

type Stage struct {
	User    *userdata.User
	Message *telegrambot.Message
}

func Begin(msg *telegrambot.Message) {
	stage := &Stage{
		User: userdata.InitUser(msg),
	}
	if stage.User.Stage == userdata.Stage_Start {
		//
	}
	log.Println(userdata.Stage_Start)
}
