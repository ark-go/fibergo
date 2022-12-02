package progDefault

import (
	"log"

	"github.com/ark-go/fibergo/internal/userdata"
)

func Volk(user *userdata.User) {
	log.Println("Второу шаг")
	g := user.Last.MapStepUser[user.GetChatUserStr()]
	g.Program = userdata.Programm_Start
	user.Last.MapStepUser[user.GetChatUserStr()] = g
	log.Println("??", user.StepUser)

	//BUG user.StepUser разобраться запутался

}
