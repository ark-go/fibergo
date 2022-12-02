package progDefault

import (
	"log"

	"github.com/ark-go/fibergo/internal/userdata"
)

// type IStep interface {
// 	Start(user *userdata.User)
// 	Next(user *userdata.User)
// }

// type Step struct{}

// func Init(user *userdata.User) *Step {
// 	st := &Step{}
// 	st.Start()
// 	return st
// }

// func (s *Step) Start() {

// }

// func (s *Step) Next(user *userdata.User) {

// }

func Start(user *userdata.User) {
	log.Println("Первый шаг")
	//	user.Last.MapStepUser[user.GetChatUserStr()] = userdata.StepUser{Stagekey: userdata.Stage_Start, Program: userdata.Programm_Volk}
	g := user.Last.MapStepUser[user.GetChatUserStr()]
	g.Program = userdata.Programm_Volk
	user.Last.MapStepUser[user.GetChatUserStr()] = g
	log.Println("??", user.StepUser)

}
