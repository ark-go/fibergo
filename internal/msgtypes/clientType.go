package msgtypes

import "github.com/nickname76/telegrambot"

type ClientType int64

const (
	Client_NotAvailable ClientType = iota
	Client_Private
	Client_Supergroup
	Client_Group
	Client_Channel
)

func (c ClientType) String() string {
	switch c {
	case Client_Private:
		return "Private"
	case Client_Supergroup:
		return "Supergroup"
	case Client_Group:
		return "Group"
	case Client_Channel:
		return "Channel"
	case Client_NotAvailable:
		return "Не определен"
	}
	return "unknown"
}

// определение типа клиента
func (rt *Info) checkClientType(msg *telegrambot.Message) {
	switch msg.Chat.Type {
	case telegrambot.ChatTypePrivate:
		rt.ClientType = Client_Private
	case telegrambot.ChatTypeChannel:
		rt.ClientType = Client_Channel
	case telegrambot.ChatTypeGroup:
		rt.ClientType = Client_Group
	case telegrambot.ChatTypeSupergroup:
		rt.ClientType = Client_Supergroup
	default:
		rt.ClientType = Client_NotAvailable
	}
}
