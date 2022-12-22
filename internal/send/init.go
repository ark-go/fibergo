package send

import (
	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

type Send struct {
	User *userdata.User
	api  *telegrambot.API
}

func Init(user *userdata.User, api *telegrambot.API) *Send {
	return &Send{
		User: user,
		api:  api,
	}

}
