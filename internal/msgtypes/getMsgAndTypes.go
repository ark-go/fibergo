package msgtypes

import (
	"errors"

	"github.com/nickname76/telegrambot"
)

type Info struct {
	// тип обновления, пришедшее сообщение Update
	UpdateType UpdateType
	// Тип клиента
	ClientType ClientType
	// Тип сообщения
	MessageType MessageType

	// Для CallbackQuery, кнопоки, необходим ID
	CallbackQueryID telegrambot.CallbackQueryID
	// Calback Data - данные из кнопки
	CalbackData string
}

// выделяем Message и тип Update
func GetMsgAndTypes(update *telegrambot.Update) (*telegrambot.Message, *Info, error) {
	rt := &Info{}
	var msg *telegrambot.Message
	// TODO доделывать остальные типы обновления
	switch {
	case update.Message != nil: // сообщение
		msg = update.Message
		rt.checkMessageType(msg)
		rt.checkClientType(msg)
		rt.UpdateType = Upd_Message
	case update.CallbackQuery != nil: // нажали кнопку inline
		i := update.CallbackQuery
		msg = i.Message
		rt.checkMessageType(msg)
		rt.CallbackQueryID = i.ID // необходим для отправки и снятия значка загрузки на кнопке
		rt.CalbackData = i.Data   // данные из кнопки
		rt.checkClientType(msg)
		rt.UpdateType = Upd_CallbackQuery
	case update.EditedMessage != nil: // отредактированное сообщение
		msg = update.EditedMessage
		rt.UpdateType = Upd_EditedMessage
	case update.EditedChannelPost != nil: // отредактированное сообщение канала
		msg = update.EditedChannelPost
		rt.UpdateType = Upd_EditedChannelPost
	case update.ChannelPost != nil: // сообщение канала
		msg = update.ChannelPost
		rt.UpdateType = Upd_ChannelPost

	default:
		return nil, nil, errors.New("не опредлен тип обновления")
	}

	return msg, rt, nil
}
