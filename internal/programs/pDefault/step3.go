package pDefault

import (
	"log"

	"github.com/ark-go/fibergo/internal/msgtypes"
	//"github.com/nickname76/telegrambot"
)

func (p *ProgDef) step3() {
	p.Send.DeleteMessageUser()
	log.Println("Прог: Это шаг 2 22222")
	if msgtypes.Msg_Text != p.Send.User.Info.MessageType {
		p.Send.SendTimeMessage("<b>Хочу только текст !!</b>пока только текст")
		return
	}
	cmd, arg := p.Program.ParseMessageCommand()
	if cmd == "" {
		log.Println("Нажмите кнопку", arg)
		p.Next("step2")
		return
	}

	log.Println("Нажмали кнопку", cmd, arg)
	//p.Send.SetCommandMenu(p.commandMenuStep2())
	p.Send.DeleteMessageUser()
	//p.Send.InlineMenuSet(p.InlineButtonStep2())
	//p.Send.SendMenu(p.MenuButton(), "Шаг-")
	p.Next("step1")
}
