/*
* Контроллер шагов
 */
package pDefault

import (
	"log"

	"github.com/ark-go/fibergo/internal/msgtypes"
	"github.com/ark-go/fibergo/internal/programs"
	"github.com/ark-go/fibergo/internal/send"
)

func init() {
	programs.AddProgram("startProg", StartProg)
}

type IProgCopy interface {
	ParseMessageCommand() (command string, args string)
	IsCommand(cmd string) bool
}
type ProgDef struct {
	*send.Send
	*programs.Program //IProgCopy
	*msgtypes.Info
}

func StartProg(program *programs.Program, step string) {
	p1 := ProgDef{
		Send:    program.Send,
		Program: program,
		Info:    program.Send.User.Info,
	}
	// удаляем сообщение которое прислал пользователь из чата
	p1.DeleteMessageUser()
	// в данной программе используем форму с Inline кнопками, поэтому только приватный режим
	if p1.ClientType != msgtypes.Client_Private {
		log.Println("Эта программа с inline кнопками, требует Private клиента, пропускаем.")
		return
	}
	// Запуск программы
	p1.StartProg(step)
}

func (p *ProgDef) StartProg(step string) {
	switch step {
	case "step1":
		p.step1()
	case "step2":
		p.step2()
	case "step3":
		p.step3()

	default:
		p.Next("step1") // поправляем запись, раз ее нет или она не правильная
		p.step1()
	}
}

// Записываем следущий шаг для клиента
func (p *ProgDef) Next(step string) {
	p.Send.User.ChangeStep(step) // запись в юзера
	p.StartProg(step)            // и переход
}
