package userdata

import (
	"log"
	"time"

	//	"github.com/ark-go/fibergo/internal/queues"
	"github.com/nickname76/telegrambot"
)

type LastUserMessage struct {
	// последнее сообщение-меню отправленное клиенту
	MessageID *telegrambot.MessageID
	ChatID    *telegrambot.ChatID
	LastTime  time.Time
}
type User struct {
	// // userId
	// UserId telegrambot.UserID
	// // chatId
	// ChatId telegrambot.ChatID
	// сообщение поступившее от пользователя
	Msg *telegrambot.Message
	// Стадия
	Stage Stagekey

	// уже был в базе
	Olden bool
	// текущее сообщение
	//	Msg *telegrambot.Message
	// тип текущего собщения
	MessageType MessageType
	// тип клиента
	ClientType ClientType
	// // Последнее сообщение от юзера
	// LastUserMessage *LastUserMessage
	// // последнее InlineMenu отправлено
	// LastInlineMenuAll LastInlineMenuAll
	// // последнее меню с кнопками отправлено
	// LastMenuAll LastMenuAll
	// // очередь отправленных сообщений
	// LastMsgQueues LastMessageQueues
	Last *Last
	// test
	Txt string
}

// Получим если можем ChatId
func (u *User) MsgChatId() (chatId telegrambot.ChatID) {
	if u.Msg != nil {
		if u.Msg.Chat != nil {
			return u.Msg.Chat.ID
		}
	}
	log.Println("Не определили ChatID")
	return
}

// Получим если сможем UserId
func (u *User) MsgUserId() (userId telegrambot.UserID) {
	if u.Msg != nil {
		if u.Msg.Chat != nil {
			return u.Msg.From.ID
		}
	}
	log.Println("Не определили ChatID")
	return
}

type Last struct {
	// Стадия
	Stage Stagekey
	// Последнее сообщение от юзера
	UserMessage *LastUserMessage
	// последнее InlineMenu отправлено
	InlineMenuAll LastInlineMenuAll
	// последнее меню с кнопками отправлено
	MenuAll LastMenuAll
	// очередь отправленных сообщений
	MsgQueues LastMessageQueues
}

func InitUser(msg *telegrambot.Message) *User {
	// if msg.From == nil || msg.Chat == nil {
	// 	log.Println("Ошибка при инициализации User")
	// 	return nil
	// }

	// структуру храним в базе для каждого пользователя
	last := &Last{
		UserMessage:   &LastUserMessage{},
		InlineMenuAll: LastInlineMenuAll{},
		MenuAll:       LastMenuAll{},
		MsgQueues:     LastMessageQueues{},
	}
	// User - что знаемпро юзера
	user := &User{
		Msg:         msg,
		Stage:       Stage_Start,
		MessageType: Msg_NotAvailable,
		ClientType:  Client_NotAvailable,
		// Last - эту структуру храним в базе, для пользователя
		Last: last,
	}
	user.ChangeMessage(msg)
	return user
}
