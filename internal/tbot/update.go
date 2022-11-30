package tbot

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	//	"time"

	"github.com/ark-go/fibergo/internal/utils"
	"github.com/nickname76/telegrambot"
)

func (b *Bot) Update(update *telegrambot.Update) {
	//msg1 := update.InlineQuery
	msg := update.Message       // чаты и группы
	inl := update.CallbackQuery // ответы кнопок inline
	chann := update.ChannelPost // каналы

	if inl != nil {
		b.inlineCommand(inl)
		return
	}
	if chann != nil {
		log.Println("канн", chann.Text)
		return
	}
	if msg == nil {
		log.Println("мсж  пусто")
		return
	}

	if msg.From.IsBot {
		// клиент есть бот
		return
	}

	if int64(msg.From.ID) == int64(msg.Chat.ID) {
		log.Println("Клиент в боте..")
	} else if int64(msg.Chat.ID) < 0 {
		log.Println("Клиент в группе..", msg.Chat.FirstName, msg.Chat.Title, msg.Chat.Type)
	}
	usr := b.Pg.LoadData(msg)
	if msg.Text == "" {
		b.sendTimeMessage(usr, "<b>Хочу только текст !!</b> сотрусь счас")
	} else {
		cmd, arg := b.ParseMessageCommand(msg)
		log.Println("Комманд", cmd, arg)
		if msg.Text == "/start" {
			b.setMenuGroupChat(usr)
		}
		// удаляем все //BUG
		b.deleteMessageUser(msg)
		b.deleteMessageMenu(usr)
		usr.Txt = time.Now().Format("02.01.2006 15:04:05") + "\n" + msg.Text
		//b.deleteMessage(msg)

		log.Println("Вошел: ", usr.TgUser.FirstName, " | ", usr.UserId, " | ", usr.LastUserMessage.LastTime.Format("02.01.2006 15:04:05"))
		// b.sendTimeMessage(usr, "Ваш последний запрос:  "+usr.LastUserMessage.LastTime.Format("02.01.2006 15:04:05")+"\r\n"+usr.Txt, 30)
		log.Println("Type mess:", usr.MessageType)
		log.Println("type client:", usr.ClientType, msg.MessageID, msg.Chat.ID)
		//BUG
		// err := b.deleteMessage(msg) // Удаляем что прислали
		// if err != nil {
		// 	log.Println("Не смогли удалить,", err.Error())
		// }
		b.sendTimeMessage(usr, fmt.Sprintf("Привет %v,  %v", msg.From.FirstName, msg.From.LastName), 10)
		// b.Api.SendMessage(&telegrambot.SendMessageParams{
		// 	ChatID:                msg.Chat.ID,
		// 	Text:                  fmt.Sprintf("Привет %v,  %v", msg.From.FirstName, msg.From.LastName),
		// 	ProtectContent:        true, // Запрет копирования - пересылки сообщения
		// 	DisableWebPagePreview: true, // отключить предпросмотр ссылок url
		// 	DisableNotification:   true, // без звука
		// })

		// if msg.Chat.Type == telegrambot.ChatTypePrivate {
		// 	b.Api.SendMessage(&telegrambot.SendMessageParams{
		// 		ChatID: msg.Chat.ID,
		// 		Text:   "<b>Вах вах вах Привате </b>",
		// 		//DisableNotification: true, // запрет копирования и пересылки
		// 		ParseMode: telegrambot.ParseModeHTML,
		// 		ReplyMarkup: &telegrambot.InlineKeyboardMarkup{
		// 			InlineKeyboard: [][]*telegrambot.InlineKeyboardButton{{
		// 				{
		// 					WebApp: &telegrambot.WebAppInfo{
		// 						URL: "https://bake.x.arkadii.ru",
		// 					},
		// 				},
		// 			}},
		// 		},
		// 	})
		// }
		//----------
		path, err := utils.GetExecPath()
		if err != nil {
			log.Println(err)
			return
		}
		path = filepath.Join(path, "img-1.jpg")
		b.sendPhoto(usr, path, 30)

		//BUG
		b.InlineMenuSet(usr, usr.CreateInlineMenu())
		//usr.Txt = msg.Text

		usr.LastUserMessage.LastTime = time.Now()
		b.deleteMessage(usr)
		b.Pg.SaveData(usr)

	}

	//b.Api.DeleteMyCommands(&telegrambot.DeleteMyCommandsParams{})
	b.Api.SetMyCommands(&telegrambot.SetMyCommandsParams{
		Commands: []*telegrambot.BotCommand{
			{
				Command:     "/start",
				Description: "Хочешь узнать?",
			},
			{
				Command:     "/bla",
				Description: "Чтото не получается?",
			},
			{
				Command:     "/bla",
				Description: "Чтото не получается?",
			}, {
				Command:     "/blahf",
				Description: "Чтото не получается опять ?",
			}, {
				Command:     "/mumu",
				Description: "Чтото не получается опять btbbttbttbb btbtb btbtbt bttbt?",
			},
		},
		// Scope: &telegrambot.BotCommandScope{
		// 	Type:   telegrambot.BotCommandScopeTypeAllChatAdministrators,
		// 	ChatID: msg.Chat.ID,
		// 	UserID: msg.From.ID,
		// },
	})

	// log.Println("mmmmmmeeeee")
	b.Api.SetChatMenuButton(&telegrambot.SetChatMenuButtonParams{
		// ChatID: msg.Chat.ID,
		MenuButton: &telegrambot.MenuButton{
			Type: telegrambot.MenuButtonTypeCommands,
			Text: "gol",
		},
	})

	// неполученные сообщения, только одно возвращает
	// updmess, _ := b.Api.GetUpdates(&telegrambot.GetUpdatesParams{
	// 	Offset:  -100,
	// 	Timeout: 0,
	// 	AllowedUpdates: *&[]telegrambot.UpdateType{
	// 		"message",
	// 	},
	// })
	// for i, v := range updmess {
	// 	log.Println(i, "------------\n", v.Message.Text)
	// }
	//b.Api.GetFile(&telegrambot.GetFileParams{})

	// type Fil telegrambot.InputFile
	// func(f *Fil) multipartFormFile() (fieldname string, filename string, reader io.Reader)

	// }

}
