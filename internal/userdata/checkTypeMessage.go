package userdata

import "github.com/nickname76/telegrambot"

type MessageType int64

const (
	Msg_NotAvailable MessageType = iota
	Msg_Text
	Msg_Video
	Msg_Caption
	Msg_Audio
	Msg_Location
	Msg_Contact
)

func (s MessageType) String() string {
	switch s {
	case Msg_Text:
		return "Text"
	case Msg_Video:
		return "Video"
	case Msg_Audio:
		return "Audio"
	case Msg_Location:
		return "Location"
	case Msg_Contact:
		return "Contact"
	case Msg_NotAvailable:
		return "Не определен"
	}
	return "unknown"
}

func (u *User) checkTypeMessage(msg *telegrambot.Message) {
	switch {
	case msg.Video != nil:
		u.MessageType = Msg_Video
	case msg.Audio != nil:
		u.MessageType = Msg_Audio
	case msg.Text != "":
		u.MessageType = Msg_Text
	case msg.Contact != nil:
		u.MessageType = Msg_Contact
	case msg.Location != nil:
		u.MessageType = Msg_Location
	default:
		u.MessageType = Msg_NotAvailable
	}
}
