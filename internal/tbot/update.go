package tbot

import (
	"log"
	"runtime"
	"time"

	"github.com/ark-go/fibergo/internal/msgtypes"
	"github.com/ark-go/fibergo/internal/programs"
	"github.com/ark-go/fibergo/internal/send"
	"github.com/ark-go/fibergo/internal/utils"
	"github.com/nickname76/telegrambot"
	"golang.org/x/exp/slices"
)

// func (b *Bot) ToIfaceBot() {}
var mStats runtime.MemStats

func (b *Bot) Update(update *telegrambot.Update) {
	runtime.ReadMemStats(&mStats)
	//	log.Printf("%+v", mStats)
	log.Printf("%+v", utils.ByteCountIEC(mStats.HeapSys-mStats.HeapAlloc))
	/*
	   !! Группа не определяет кто нажимает Inline кнопки.. они идут от группы
	   !! и сообщение в группе идет всем в группу
	   !! такчто надо отключить работу с картинками, по юзерам в группе
	   !! текстовые сообщения идентифицируются нормально
	*/
	//msg1 := update.InlineQuery
	log.Println(">---------- пришло --------------------<")

	msg, rt, err := msgtypes.GetMsgAndTypes(update)
	if err != nil {
		log.Println("Ошибка типа", err.Error())
		log.Panicln(" Типы не забыть !!!!", rt) // nil отуда?
		//! не ставлю возврат, надо найти какой Update все Update надо вписать  return
	}

	/*
		!! боты отправляют, нажатие inline-кнопки тоже bot!
		 если это бот отправляет чтото кроме нажатия кнопки, пропускаем
	*/
	// ввод из канала
	if rt.UpdateType == msgtypes.Upd_ChannelPost { // .Upd_EditedChannelPost {
		log.Println("Канал пропускаем, пока")
		return
	}
	//!! редактирование канала
	if rt.UpdateType == msgtypes.Upd_EditedChannelPost { // .Upd_EditedChannelPost {
		log.Println("Канал редактируем  - пропускаем, пока")
		return
	}
	//!! в канале нет From
	if msg.From != nil && msg.From.IsBot { // это бот
		// нажатие кнопок проискодит от имени бота
		// если это бот но не нажимает кнопки то пропускаем
		// не обслуживаем бота , кроме нажатия кнопок
		if rt.UpdateType != msgtypes.Upd_CallbackQuery {
			log.Println("Это бот, не общаемся с ними")
			return
		}
	}

	isDeleteMessage := false
	if msg.IsAutomaticForward {
		// сообщенее переслано из канала в связанную группу
		// в которой у нас есть бот мы не должны реагировать на это сообщение,
		// и не обслуживать его, оно так и останется в группе
		//! надо бы пропустить и отменять у разбора в ProgramBegin
		log.Println("Из канала пришло в группу, пропускаем")
		return
	}
	switch rt.UpdateType {
	case msgtypes.Upd_EditedMessage:
		// зачем нам редактировать сообщения, удалим целиком раз не нравится
		isDeleteMessage = true
	case msgtypes.Upd_ChannelPost,
		msgtypes.Upd_EditedChannelPost:
		// не обращаем внимание на сообщения из канала
		return
	default:
		//return
	}

	if msg.Chat.ID < 0 {
		// BUG Чтото надо тут делть чтоб определить группу
		return
	}

	if isDeleteMessage {
		// удаляем сообщение пользователя, если нам не нравится его тип
		err := b.Api.DeleteMessage(&telegrambot.DeleteMessageParams{
			ChatID:    msg.Chat.ID,
			MessageID: msg.MessageID,
		})
		if err != nil {
			log.Println("Не удалить сообщение", err.Error())
		}
		return
	}

	/*
		switch {
		// 		log.Println("callback:", update.CallbackQuery.ChatInstance, msg.Text)
		case update.MyChatMember != nil: // чат заблокировали или разюлокировали
		case update.ChatMember != nil: // участник чата обновил свой статус
		case update.ChatJoinRequest != nil: // Запрос на вступление в чат
		case update.Poll != nil: // опрос завершен или изменилось состояние, только опросы бота
		case update.PollAnswer != nil: // пользователь изменил свой ответ в опросе, только опросы бота
		default:
			log.Println("Нет отлавливаемого сообщения")
		}
	*/

	// загрузим данные из базы, если там был пользователь
	// возвращаем User
	user, err := b.Pg.LoadData(msg)
	if err != nil {
		log.Println(err.Error()) // нет From e MSJ
		return
	}
	log.Println(">--->->user.info", user.Info)
	user.Info = rt
	log.Println(">--->->user.info", user.Info)
	// !! Необходимо разместить код авторизации как программу по умолчанию при отсутстви  данных в базе
	var svoi = []int64{
		266848998,  //ark
		1624458545, // alina
		//	-1001492038864, // группа
	}
	curUser := user.GetUserId()
	if slices.Index(svoi, int64(curUser)) == -1 {
		b.Api.SendMessage(&telegrambot.SendMessageParams{ChatID: user.GetChatId(), Text: "❌<b>Hi!</b>"})
		return
	}
	// !! ^^^

	//!  надо вернуть u.ChangeMessage(msg) /
	// Программ подготовка
	//user.Bot = b
	send := send.Init(user, b.Api) // передаем api
	prog := programs.Init(send)    // передаем юзера в программ - структуру

	//send.DeleteMessageUser()

	log.Println("Тип пришел: ", msg.MessageID, user.Info.MessageType, "Client Type: ", rt.ClientType)
	if user.Info.ClientType == msgtypes.Client_Group {
		log.Println("Группа: ", user.Info.ClientType, msg.Chat.FirstName, msg.Chat.Title, msg.Chat.Type)
	}

	log.Println("Вошел: ", user.Info.UpdateType, "| ", user.Msg.Chat.ID, "|", user.GetUserId(), " | ", user.Msg.From.Username, "|", user.UserData.UserMessage.LastTime.Format("02.01.2006 15:04:05"))
	/*
	 !! тест Отправляем кучку фоток
	*/
	// send.SendGroupPhoto([]string{"/home/arkadii/ProjectsGo/fibergo/bin/bot/img-2.jpg",
	// 	"/home/arkadii/ProjectsGo/fibergo/bin/bot/i2mg-2.jpg",
	// 	"/home/arkadii/ProjectsGo/fibergo/bin/bot/img-2.jpg",
	// })

	//!! По кнопке приходит само сообщение. если там у кнопки была картинка то картинка и приходит
	if user.Info.UpdateType != msgtypes.Upd_CallbackQuery {
		if user.Info.MessageType == msgtypes.Msg_Photo {
			log.Println("photo", len(msg.Photo))
			send.LoadFileFromFileID(msg.Photo[len(msg.Photo)-1].FileID)
		}

		if user.Info.MessageType == msgtypes.Msg_Document {
			log.Println("Докум")
			send.LoadFileFromFileID(msg.Document.FileID)
		}
	}
	// удаляет последние сообщения из чата пользователя
	// указываем кол-во последних
	// send.DeleteLastMessages(msg, 10)
	// запустим горутину ?
	go func() {
		prog.ProgramBegin()
		//time.Sleep(1 * time.Second)
		//log.Println("time>", user.Msg.Text)
		// send.DeleteMessage()
		prog.User.UserData.UserMessage.LastTime = time.Now()
		//log.Println("inl25:", *user.UserData.InlineMenuAll[user.GetChatId()].MessageID)
		b.Pg.SaveData(user)

	}()

}
