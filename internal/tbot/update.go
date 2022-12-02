package tbot

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ark-go/fibergo/internal/programs"
	"github.com/ark-go/fibergo/internal/stages"
	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

func (b *Bot) Update(update *telegrambot.Update) {
	//msg1 := update.InlineQuery
	log.Println(">---------- пришло --------------------<")
	usr := &userdata.User{}
	var msg *telegrambot.Message
	if msg = update.Message; msg != nil { // чаты и группы
		usr = b.Pg.LoadData(msg)
		if msg.IsAutomaticForward {
			log.Println("Сообщение переслано из канала в чат, а из чата сюда") //, msg.ForwardFrom.FirstName)
		}
	}
	if usr.ClientType == userdata.Client_Channel {
		log.Println("Канал - пропускаем не обслуживаем")
		// TODO разобраться с каналами
		return
	}

	// else if msg = update.ChannelPost; msg != nil {
	// 	usr = b.Pg.LoadData(msg)
	// }
	// ответы кнопок inline
	if inl := update.CallbackQuery; inl != nil {
		b.inlineCommand(inl)
		return
	}
	if msg == nil {
		log.Println("Пока пропустим  - пустое сообщение вероятно канал")
		return
	}
	//usr.Last.MapStepUser[usr.GetChatUserStr()] = usr.StepUser
	log.Printf("Программа: %+v", usr.Last.MapStepUser[usr.GetChatUserStr()])
	//usr.Last.MapStepUser[usr.GetChatUserStr()] = userdata.StepUser{Stagekey: userdata.Stage_Start, Program: userdata.Programm_Volk}

	log.Println("Message Type: ", usr.MessageType)
	log.Println("Client Type: ", usr.ClientType, usr.GetChatUserStr())
	if usr.ClientType == userdata.Client_Group {
		log.Println("Группа: ", usr.ClientType, msg.Chat.FirstName, msg.Chat.Title, msg.Chat.Type)
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
	}
	//  else if int64(msg.Chat.ID) < 0 {
	// 	log.Println("Клиент в группе..", msg.Chat.FirstName, msg.Chat.Title, msg.Chat.Type)
	// }

	//	usr := b.Pg.LoadData(msg)

	if msg.Text == "" {
		b.sendTimeMessage(usr, "<b>Хочу только текст !!</b> сотрусь счас")
		b.deleteMessageUser(msg)
	} else {
		cmd, arg := b.ParseMessageCommand(msg)
		log.Println("Комманд", cmd, arg)
		if msg.Text == "/start" {
			b.MenuSend(usr, "Старт:")
		} else {
			b.deleteMessageMenu(usr)
		}

		b.deleteMessageUser(msg) // удаление всего - убрать?

		stages.Begin(msg)
		usr.Txt = time.Now().Format("02.01.2006 15:04:05") + "\n" + msg.Text

		log.Println("Вошел: ", "| ", usr.MsgUserId(), " | ", usr.Last.UserMessage.LastTime.Format("02.01.2006 15:04:05"))
		//BUG
		b.sendTimeMessage(usr, fmt.Sprintf("Привет %v  %v %s", msg.From.FirstName, msg.From.LastName, strings.Repeat("ᅠ", 20)), 10)
		programs.StartProgram(usr)
		//BUG
		b.InlineMenuSet(usr, usr.CreateInlineMenu())
		usr.Last.UserMessage.LastTime = time.Now()
		b.deleteMessage(usr)
		b.Pg.SaveData(usr)

	}

	// b.Api.DeleteMyCommands(&telegrambot.DeleteMyCommandsParams{
	// 	Scope: &telegrambot.BotCommandScope{
	// 		Type: telegrambot.BotCommandScopeTypeAllChatAdministrators, // в чаты где бот адмистратором
	// 		//Type:   telegrambot.BotCommandScopeTypeChat, // в чат откуда вылез пользователь
	// 		ChatID: msg.Chat.ID,
	// 		UserID: msg.From.ID,
	// 	},
	// })
	b.Api.SetMyCommands(&telegrambot.SetMyCommandsParams{
		Commands: []*telegrambot.BotCommand{
			{
				Command:     "/start",
				Description: "Хочешь узнать?",
			},
			{
				Command:     "/start2",
				Description: "Чтото не получается?",
			},
		},
		Scope: &telegrambot.BotCommandScope{
			//	Type:   telegrambot.BotCommandScopeTypeAllChatAdministrators, // в чаты где бот адмистратором
			Type:   telegrambot.BotCommandScopeTypeChat, // в чат откуда вылез пользователь
			ChatID: msg.Chat.ID,
			UserID: msg.From.ID,
		},
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
