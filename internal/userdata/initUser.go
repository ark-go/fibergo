package userdata

import (
	"log"
	"time"

	//	"github.com/ark-go/fibergo/internal/queues"
	"github.com/nickname76/telegrambot"
)

// type LastInlineMenu struct {
// 	// последнее сообщение-меню отправленное клиенту
// 	MessageID *telegrambot.MessageID
// 	ChatID    *telegrambot.ChatID
// 	LastTime  time.Time
// }
// type LastInlineMenuAll []LastInlineMenuAll

type LastUserMessage struct {
	// последнее сообщение-меню отправленное клиенту
	MessageID *telegrambot.MessageID
	ChatID    *telegrambot.ChatID
	LastTime  time.Time
}
type User struct {
	// userId
	UserId telegrambot.UserID
	// chatId
	ChatId telegrambot.ChatID
	// сообщение поступившее от пользователя
	//LastUsrMessage telegrambot.MessageID
	// Стадия
	Stage string
	// telegram User
	TgUser *telegrambot.User
	// уже был в базе
	Olden bool
	// текущее сообщение
	//	Msg *telegrambot.Message
	// тип текущего собщения
	MessageType MessageType
	// тип клиента
	ClientType ClientType
	// Последнее сообщение от юзера
	LastUserMessage *LastUserMessage
	// последнее InlineMenu отправлено
	//LastInlineMenu *LastInlineMenu
	LastInlineMenuAll LastInlineMenuAll
	// последнее меню с кнопками отправлено
	//LastMenu    *LastMenu
	LastMenuAll LastMenuAll
	// очередь отправленных сообщений
	LastMsgQueues LastMessageQueues

	// test
	Txt string
}

func InitUser(msg *telegrambot.Message) *User {
	if msg.From == nil || msg.Chat == nil {
		log.Println("Ошибка при инициализации User")
		return nil
	}
	user := &User{
		MessageType: Msg_NotAvailable,
		ClientType:  Client_NotAvailable,
		//LastUsrMessage:  msg.MessageID,
		LastUserMessage: &LastUserMessage{},
		//LastInlineMenu:  &LastInlineMenu{},
		//LastMenu:        &LastMenu{},
		LastInlineMenuAll: LastInlineMenuAll{},
		LastMenuAll:       LastMenuAll{},
		LastMsgQueues:     LastMessageQueues{},
	}
	user.ChangeMessage(msg)
	return user
}
