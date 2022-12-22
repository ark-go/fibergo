package msgtypes

import "github.com/nickname76/telegrambot"

type MessageType int64

const (
	Msg_NotAvailable MessageType = iota
	Msg_Text
	Msg_Video
	Msg_Photo
	Msg_Caption
	Msg_Document
	Msg_Audio
	Msg_Location
	Msg_Contact
	Msg_Voice
	// сообщение в комментарии к каналу, в связанной группе
	Msg_ReplyToMessage
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
	case Msg_Document:
		return "Document"
	case Msg_Voice:
		return "Voice"
	case Msg_ReplyToMessage:
		return "ReplyToMessage"
	case Msg_NotAvailable:
		return "Не определен"
	case Msg_ArkReserved:
		return "Не реализовано"
	}
	return "unknown"
}

// Определение типа сообщения
func (rt *Info) checkMessageType(msg *telegrambot.Message) {
	switch {
	case msg.Animation != nil:
	case msg.Audio != nil:
		rt.MessageType = Msg_Audio
	case msg.Contact != nil:
		rt.MessageType = Msg_Contact
	case msg.Dice != nil:
		rt.MessageType = Msg_ArkReserved
	case msg.Document != nil:
		rt.MessageType = Msg_Document
	case msg.Game != nil:
		rt.MessageType = Msg_ArkReserved
	case msg.Invoice != nil: // счет на оплату
		rt.MessageType = Msg_ArkReserved
	case msg.Location != nil:
		rt.MessageType = Msg_Location
	case msg.Photo != nil:
		rt.MessageType = Msg_Photo
	case msg.PinnedMessage != nil:
		rt.MessageType = Msg_ArkReserved
	case msg.Poll != nil:
		rt.MessageType = Msg_ArkReserved
	case msg.ProximityAlertTriggered != nil: //активировал оповещение о приближении другого пользователя
		rt.MessageType = Msg_ArkReserved
	case msg.ReplyToMessage != nil: //Для ответов исходное сообщение. Обратите внимание, что объект Message в этом поле не будет содержать дополнительных полей answer_to_message, даже если он сам является ответом.
		rt.MessageType = Msg_ReplyToMessage
	case msg.Sticker != nil:
		rt.MessageType = Msg_ArkReserved
	case msg.SuccessfulPayment != nil:
		rt.MessageType = Msg_ArkReserved
	case msg.Text != "": // Текст
		rt.MessageType = Msg_Text
	case msg.Venue != nil: // место проведения, информация о месте проведения
		rt.MessageType = Msg_ArkReserved
	case msg.Video != nil: // видео
		rt.MessageType = Msg_Video
	case msg.VideoChatScheduled != nil: // видеочат запланирован
		rt.MessageType = Msg_ArkReserved
	case msg.VideoChatStarted != nil: // видеочат запущен
		rt.MessageType = Msg_ArkReserved
	case msg.VideoChatEnded != nil: // видеочат остановлен
		rt.MessageType = Msg_ArkReserved
	case msg.Voice != nil: // голосовое сообщение
		rt.MessageType = Msg_Voice
	default:
		rt.MessageType = Msg_NotAvailable
	}
}
