package pDefault

import (
	"errors"
	"fmt"
	"log"

	"github.com/ark-go/fibergo/internal/msgtypes"
	"github.com/ark-go/fibergo/internal/nasa"
	"github.com/nickname76/telegrambot"
)

// первый шаг ожидаем  текст
// выкатим картинку, зададим второй шаг и выйдем
func (p *ProgDef) step1() {
	log.Println("Прог: Это шаг 1 11111")
	//	log.Println("inl259:", *p.Send.User.UserData.InlineMenuAll[p.Send.User.GetChatId()].MessageID)
	if !p.IsCommand("start") && p.UpdateType != msgtypes.Upd_CallbackQuery {
		p.DeleteMessageUser()
		p.AnswerCallbackQuery("")
		p.SendTimeMessage("я первый Нажмте кнопку !!")
		return
	}
	str := fmt.Sprintf("\n<b>➖➖➖</b><pre>%s</pre>", p.CalbackData)
	if msg, err := p.MenuCaption(p.menuButton1(), "Прог 1:"+str); err != nil {

		if errors.Is(err, msgtypes.ErrNotFound) || errors.Is(err, msgtypes.ErrNotFoundKeyUser) {
			p.Send.MenuPhoto(p.menuButton1(), "img-210.png", p.User.UserData.InlineMenuAll.GetFileId(p.User.GetChatId()), "Шаг 1")
		} else if !errors.Is(err, msgtypes.ErrNotModified) {
			log.Println("Замена не прошла", msg, err.Error())
		}
	} else {

		log.Println("Замена", msg) //! нет msg ,,?
	}

	// меню, если пусто там будет только /start
	p.SetCommandMenu(nil)
	log.Println("Установили Меню")
	switch p.CalbackData {
	case "stepTwo":
		p.Next("step2")
	case "nasa":
		p.InlineMenuDelete()
		p.ChatAction()
		nas, err := nasa.LoadNasa()
		if err != nil {
			log.Println("nasa error", err)
			return
		}
		p.Photo(nil, nas[0].Url, "", nas[0].Title)
		p.MenuPhoto(p.menuButton1(), "img-210.png", p.User.UserData.InlineMenuAll.GetFileId(p.User.GetChatId()), "", "")
	case "epic":
		p.InlineMenuDelete()
		p.ChatAction()
		epic, err := nasa.LoadEpic()
		if err != nil || epic == nil || len(epic) < 1 {
			p.SendTimeMessage("Попробуй еще раз, не оказалось снимков на дату \nили Menu -> сброс", 5)
			return
		} else {
			log.Println("картинок всего:", len(epic))
			var arrUrlCap = make([][]string, len(epic))
			for i, val := range epic {
				arrUrlCap[i] = []string{
					val.Url,
					val.Caption,
				}
			}

			if err := p.PhotoGroup(nil, arrUrlCap); err != nil {
				p.SendTimeMessage("<b>Произошла ошибка при получении картинок..</b>", 5)
			}
			p.MenuPhoto(p.menuButton1(), "img-210.png", p.User.UserData.InlineMenuAll.GetFileId(p.User.GetChatId()), "", "")
		}
	case "fox":
		url, err := nasa.Fox()
		p.SendPhotoFromArrayStr(url, err, "fox")

	case "dog":
		url, err := nasa.Dog()
		p.SendPhotoFromArrayStr(url, err, "dog")

	case "cat":
		url, err := nasa.Cat()
		p.SendPhotoFromArrayStr(url, err, "cat")

	case "duck":
		url, err := nasa.Duck()
		p.SendPhotoFromArrayStr(url, err, "duck")

	case "birds":
		url, err := nasa.Birds()
		p.SendPhotoFromArrayStr(url, err, "birds")

	case "anime1":
		url, err := nasa.Anime1()
		p.SendPhotoFromArrayStr(url, err, "anime1")

	case "anime5":
		url, err := nasa.Anime5()
		p.SendPhotoFromArrayStr(url, err, "anime5")

	case "anime4":
		url, err := nasa.Anime4()
		p.SendPhotoFromArrayStr(url, err, "anime4")
	case "reset":
		p.DeleteMessage()
	}
	// ответ чтобы снять значеок часов - ожидания на кнопке
	p.AnswerCallbackQuery("")

}

func (p *ProgDef) IfErrorMessage(err error, nameErr string) bool {
	// var logm string
	// if len(logmess) > 0 {
	// 	logm = "Error " + strings.Join(logmess, "")
	// }
	if err != nil {
		log.Println("Error "+nameErr, err)
		p.SendTimeMessage("Попробуй еще раз, не оказалось снимков \nили Menu -> сброс", 5)
		return true
	}
	return false
}

func (p *ProgDef) SendPhotoFromArrayStr(url []string, err error, nameErr string) {
	p.InlineMenuDelete()
	p.ChatAction()
	//url, err := nasa.Anime4()
	if p.IfErrorMessage(err, nameErr) {
		return
	}
	p.Photo(nil, url[0], "")
	p.MenuPhoto(p.menuButton1(), "img-210.png", p.User.UserData.InlineMenuAll.GetFileId(p.User.GetChatId()), "", "")
}

func (p *ProgDef) menuButton1() *telegrambot.InlineKeyboardMarkup {
	return &telegrambot.InlineKeyboardMarkup{
		InlineKeyboard: [][]*telegrambot.InlineKeyboardButton{
			// 	{
			// 	{
			// 		Text:         "Информация",
			// 		CallbackData: "info@start",
			// 	},
			// },
			{
				{
					Text:         "Туда",
					CallbackData: "stepTwo",
				},

				{
					Text:         "🪐",
					CallbackData: "nasa",
				},
				{
					Text:         "🌎",
					CallbackData: "epic",
				},
				{
					Text:         "🦊",
					CallbackData: "fox",
				},
				{
					Text:         "🐶",
					CallbackData: "dog",
				},
				{
					Text:         "🐱",
					CallbackData: "cat",
				},
			}, {
				{
					Text:         "🦆",
					CallbackData: "duck",
				},
				{
					Text:         "🦅",
					CallbackData: "birds",
				},
				{
					Text:         "👯",
					CallbackData: "anime1",
				},
				{
					Text:         "👯",
					CallbackData: "anime5",
				},
				{
					Text:         "🤪",
					CallbackData: "anime4",
				},
				{
					Text:         "Сброс",
					CallbackData: "reset",
				},
			}},
	}
}
