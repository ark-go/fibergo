package userdata

import (
	"log"
	"strconv"
	"time"

	"github.com/ark-go/fibergo/internal/msgtypes"
	"github.com/nickname76/telegrambot"
)

type LastUserMessage struct {
	// последнее сообщение-меню отправленное клиенту
	MessageID *telegrambot.MessageID
	ChatID    *telegrambot.ChatID
	LastTime  time.Time
}

type User struct {
	// Текущий уникальный ID  - ChatId + UserId
	Uid string
	// сообщение поступившее от пользователя
	Msg *telegrambot.Message
	// уже был в базе
	Olden bool
	// информация о входящем сообщении
	Info *msgtypes.Info
	// структура для хранения в базе информация пользователя
	UserData *UserData
	// test
	Txt string
}

// Получим если можем ChatId или 0
func (u *User) GetChatId() (chatId telegrambot.ChatID) {
	if u.Msg != nil {
		if u.Msg.Chat != nil {
			return u.Msg.Chat.ID
		}
	}
	log.Println("Не определили ChatID")
	return
}

// Получим если сможем UserId
func (u *User) GetUserId() (userId telegrambot.UserID) {
	if u.Msg != nil {
		if u.Msg.From != nil && !u.Msg.From.IsBot {
			return u.Msg.From.ID
		} else {
			if u.Msg.Chat != nil {
				return telegrambot.UserID(u.Msg.Chat.ID)
			}
		}
	}
	log.Println("Не определили UserID")
	return
}

// Выдаст текущий ID идентификатор получающийся
//
//	сложением Id Chat и ID User
func (u *User) GetUid() string {
	ch := u.GetChatId()
	us := u.GetUserId()
	u.Uid = strconv.FormatInt(int64(ch), 10) + ":" + strconv.FormatInt(int64(us), 10)
	return u.Uid // ! ??
}

// структуру UserData храним в базе для каждого пользователя
type UserData struct {
	// Последнее сообщение от юзера
	UserMessage *LastUserMessage
	// последнее InlineMenu отправлено
	InlineMenuAll LastInlineMenuAll
	// последнее меню с кнопками отправлено
	MenuAll LastMenuAll
	// очередь отправленных сообщений
	MsgQueues LastMessageQueues
	// Текущая программа - шаг
	NextStep NextStep
}

func InitUser(msg *telegrambot.Message) *User {

	// структуру UserData храним в базе для каждого пользователя
	UserData := &UserData{
		UserMessage:   &LastUserMessage{},
		InlineMenuAll: LastInlineMenuAll{},
		MenuAll:       LastMenuAll{},
		MsgQueues:     LastMessageQueues{},
	}

	// User - что знаем про юзера
	user := &User{
		Msg: msg,
		// UserData - эту структуру храним в базе, для пользователя
		UserData: UserData,
	}
	return user
}
