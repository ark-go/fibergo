package userdata

import "github.com/nickname76/telegrambot"

type MessageType int64

const (
	Msg_NotAvailable MessageType = iota
	Msg_Text
	Msg_Video
	Msg_Photo
	Msg_Caption
	Msg_Audio
	Msg_Location
	Msg_Contact
	Msg_Voice
	Msg_ArkReserved
)

func (s MessageType) String() string {
	switch s {
	case Msg_Text:
		return "Text"
	case Msg_Video:
		return "Video"
	case Msg_Photo:
		return "Photo"
	case Msg_Audio:
		return "Audio"
	case Msg_Location:
		return "Location"
	case Msg_Contact:
		return "Contact"
	case Msg_Voice:
		return "Voice"
	case Msg_NotAvailable:
		return "Не определен"
	case Msg_ArkReserved:
		return "Не реализовано"
	}
	return "unknown"
}

func (u *User) checkMessageType(msg *telegrambot.Message) {
	switch {
	case msg.Animation != nil:
	case msg.Audio != nil:
		u.MessageType = Msg_Audio
	case msg.Contact != nil:
		u.MessageType = Msg_Contact
	case msg.Dice != nil:
		u.MessageType = Msg_ArkReserved
	case msg.Document != nil:
		u.MessageType = Msg_ArkReserved
	case msg.Game != nil:
		u.MessageType = Msg_ArkReserved
	case msg.Invoice != nil: // счет на оплату
		u.MessageType = Msg_ArkReserved
	case msg.Location != nil:
		u.MessageType = Msg_Location
	case msg.Photo != nil:
		u.MessageType = Msg_ArkReserved
	case msg.PinnedMessage != nil:
		u.MessageType = Msg_ArkReserved
	case msg.Poll != nil:
		u.MessageType = Msg_ArkReserved
	case msg.ProximityAlertTriggered != nil: //активировал оповещение о приближении другого пользователя
		u.MessageType = Msg_ArkReserved
	case msg.ReplyToMessage != nil: //Для ответов исходное сообщение. Обратите внимание, что объект Message в этом поле не будет содержать дополнительных полей answer_to_message, даже если он сам является ответом.
		u.MessageType = Msg_ArkReserved
	case msg.Sticker != nil:
		u.MessageType = Msg_ArkReserved
	case msg.SuccessfulPayment != nil:
		u.MessageType = Msg_ArkReserved
	case msg.Text != "": // Текст
		u.MessageType = Msg_Text
	case msg.Venue != nil: // место проведения, информация о месте проведения
		u.MessageType = Msg_ArkReserved
	case msg.Video != nil: // видео
		u.MessageType = Msg_Video
	case msg.VideoChatScheduled != nil: // видеочат запланирован
		u.MessageType = Msg_ArkReserved
	case msg.VideoChatStarted != nil: // видеочат запущен
		u.MessageType = Msg_ArkReserved
	case msg.VideoChatEnded != nil: // видеочат остановлен
		u.MessageType = Msg_ArkReserved
	case msg.Voice != nil: // голосовое сообщение
		u.MessageType = Msg_Voice
	default:
		u.MessageType = Msg_NotAvailable
	}
}
