package pDefault

import (
	"errors"
	"fmt"
	"log"

	"github.com/ark-go/fibergo/internal/msgtypes"
	"github.com/nickname76/telegrambot"
)

func (p *ProgDef) step2() {
	log.Println("Прог: Это шаг 2 2 2 2 ")
	// ожидаем нажатие inline кнопки
	if msgtypes.Upd_CallbackQuery != p.UpdateType {
		p.DeleteMessageUser()
		p.SendTimeMessage("я второй Нажмте кнопку !!")

		return
	}
	str := fmt.Sprintf("\n<b>➖➖➖</b><pre>%s</pre>", p.CalbackData)
	if msg, err := p.MenuCaption(p.menuButton2(), "Прог 2:"+str); err != nil {
		if errors.Is(err, msgtypes.ErrNotFound) {
			p.Send.MenuPhoto(p.menuButton2(), "img-210.png", "Шаг 2")
		} else if !errors.Is(err, msgtypes.ErrNotModified) {
			log.Println("Замена не прошла", msg, err.Error())
		}
	} else {
		log.Println("Замена 2", msg) //! нет msg ,,?
	}
	// установка текстового меню у пользователя или в группе в строке редактирования
	p.SetCommandMenu(p.commandMenuStep2())
	// нажатия кнопок проверим
	switch p.CalbackData {
	case "stepOne":
		p.Next("step1")
	}
	// снимаем значок ожидания у кнопки
	p.AnswerCallbackQuery("")
	//log.Println("кнопка:", p.Send.User.Info.CalbackData)

}

// func menu(replyMark *telegrambot.InlineKeyboardMarkup){

// }

func (p *ProgDef) commandMenuStep2() (botCommand []*telegrambot.BotCommand) {
	botCommand = []*telegrambot.BotCommand{
		{
			Command:     "/start",
			Description: "Старт или сброс",
		},
		{
			Command:     "/go",
			Description: "Иди гуляй",
		},
	}
	return
}

func (p *ProgDef) menuButton2() *telegrambot.InlineKeyboardMarkup {
	return &telegrambot.InlineKeyboardMarkup{
		InlineKeyboard: [][]*telegrambot.InlineKeyboardButton{{
			{
				Text:         "Назад",
				CallbackData: "info@start",
			},
		}, {
			{
				Text:         "11",
				CallbackData: "Я вася@петров",
			},
			{
				Text:         "Фотки",
				CallbackData: "stepOne",
			},
			{
				Text:         "33",
				CallbackData: "2334",
			},
		}},
	}
}
