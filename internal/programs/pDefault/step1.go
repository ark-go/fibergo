package pDefault

import (
	"errors"
	"fmt"
	"log"

	"github.com/ark-go/fibergo/internal/msgtypes"
	"github.com/nickname76/telegrambot"
)

// первый шаг ожидаем  текст
// выкатим картинку, зададим второй шаг и выйдем
func (p *ProgDef) step1() {
	log.Println("Прог: Это шаг 1 11111")
	//	log.Println("inl259:", *p.Send.User.UserData.InlineMenuAll[p.Send.User.GetChatId()].MessageID)
	if !p.Program.IsCommand("start") && msgtypes.Upd_CallbackQuery != p.Send.User.Info.UpdateType {
		p.Send.DeleteMessageUser()
		p.Send.AnswerCallbackQuery("")
		p.Send.SendTimeMessage("я первый Нажмте кнопку !!")
		return
	}
	str := fmt.Sprintf("\n<b>➖➖➖</b><pre>%s</pre>", p.Send.User.Info.CalbackData)
	if msg, err := p.Send.EditMessageCaption(p.InlineButtonStep1(), "Прог1:"+str); err != nil {

		if errors.Is(err, msgtypes.ErrNotFound) || errors.Is(err, msgtypes.ErrNotFoundKeyUser) {
			p.Send.InlineMenuSet(p.InlineButtonStep1(), "Шаг 1")
		} else if !errors.Is(err, msgtypes.ErrNotModified) {
			log.Println("Замена не прошла", msg, err.Error())
		}
	} else {

		log.Println("Замена", msg) //! нет msg ,,?
	}

	// меню, если пусто там будет только /start
	p.Send.SetCommandMenu(nil)
	log.Println("Установили Меню")
	switch p.Send.User.Info.CalbackData {
	case "stepTwo":
		p.Next("step2")
	}
	// ответ чтобы снять значеок часов - ожидания на кнопке
	p.Send.AnswerCallbackQuery("")

}

func (p *ProgDef) InlineButtonStep1() *telegrambot.InlineKeyboardMarkup {
	return &telegrambot.InlineKeyboardMarkup{
		InlineKeyboard: [][]*telegrambot.InlineKeyboardButton{{
			{
				Text:         "Информация",
				CallbackData: "info@start",
			},
		}, {
			{
				Text:         "Шаг2",
				CallbackData: "stepTwo",
			},
			{
				Text:         "22",
				CallbackData: "23",
			},
			{
				Text:         "33",
				CallbackData: "2334",
			},
			{
				Text:         "44",
				CallbackData: "2334",
			},
			{
				Text:         "55",
				CallbackData: "Сброс",
			},
		}},
	}
}
