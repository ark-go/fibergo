package msgtypes

type UpdateTypeInt int

const (
	Upd_notavailable UpdateTypeInt = iota
	Upd_Message
	Upd_EditedMessage
	Upd_ChannelPost
	Upd_EditedChannelPost
	Upd_CallbackQuery
	Upd_MyChatMember
	Upd_ChatJoinRequest
	Upd_Pool
	Upd_PoolAnswer
)
