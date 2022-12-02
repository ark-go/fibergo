package programs

import (
	"github.com/ark-go/fibergo/internal/programs/progDefault"
	"github.com/ark-go/fibergo/internal/userdata"
)

type prog map[userdata.Program]func(user *userdata.User)

var Prog prog = prog{
	userdata.Programm_Start: progDefault.Start,
	userdata.Programm_Volk:  progDefault.Volk,
}

func StartProgram(user *userdata.User) {
	step := user.Last.MapStepUser[user.GetChatUserStr()]
	// switch step.Program {
	// case userdata.Programm_Start:
	// 	log.Println("Программа Старт программы")
	// case userdata.Programm_Volk:
	// 	log.Println("Программа Волки ?")
	// default:
	// }

	if v, ok := Prog[step.Program]; ok {
		v(user)
	}

}
