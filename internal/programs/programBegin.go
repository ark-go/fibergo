package programs

import (
	"log"

	"github.com/ark-go/fibergo/internal/send"
)

type Program struct {
	*send.Send
}

// map type  program/func
type programList map[string]func(*Program, string)

// карта для хранения пользовательских программ,
// заполняется пользовательским кодом,
// где указывается точка входа программы и шага в программе
var ProgramList = make(map[string]func(*Program, string))

// передача Send в контроллер
func Init(send *send.Send) *Program {
	return &Program{
		Send: send,
	}
}

// Добавление программы
func AddProgram(progname string, fn func(*Program, string)) {
	ProgramList[progname] = fn
}

// Контроллер выбора программы
func (p *Program) ProgramBegin() {
	// читаем программу и шаг если сохранена у пользователя
	// или создаем стартовую
	prog, step := p.checkOrInitProgram()
	// проверим была ли команда /start
	command, arg := p.ParseMessageCommand()
	log.Println("Команда:", command, arg)
	// Проверка команды до выбора программы
	switch command {
	// если команда старт - сбрасываем программу
	case "start":
		prog, step = p.checkOrInitProgram(true)

	}
	// запуск текущей программы
	if f, ok := ProgramList[prog]; ok {
		f(p, step)
	} else {
		log.Println("нет программы", prog, ProgramList)
	}
}

/*
	Сброс/установка или считывание программы с шагом с клиента

если у клиента нет прогаммы или шага - устанавливается первая программа и шаг 1
если есть - возвращается программа и шаг из клиента, которая была вероятно, в базе
*/
func (p *Program) checkOrInitProgram(reset ...bool) (prog string, step string) {
	// попытка получить програму у пользователя
	prog, step = p.Send.User.GetStep()
	// если нет програмы или устанвлен флаг сброса
	if prog == "" || (len(reset) > 0 && reset[0]) {
		log.Println("сброс программы")
		// инициализация программы если не была задана команда
		// или команда /start
		// записываем данные в клиент
		prog := p.Send.User.ChangeProgram("startProg")
		step := p.Send.User.ChangeStep("step1")
		// удаляем сообщение из чата если там оно есть
		p.DeleteMessage()
		p.Send.InlineMenuDelete()
		// сбрасывыем запись о форме-картинке меню с кнопками.
		p.Send.User.UserData.InlineMenuAll.Drop(p.Send.User.GetChatId())
		return prog, step
	}

	if step == "" {
		// если нет шага или  prog пустой
		step = p.Send.User.ChangeStep("step1")

	}
	return prog, step
}

/*
	переход на следущий шаг

шаг string , шаг должен существовать в программе
*/
func (p *Program) Next(step string) {
	p.Send.User.ChangeStep(step)
}
