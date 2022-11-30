package userdata

import "time"

type LastMessage struct {
	MessageID  *int64
	ChatID     *int64
	TimeDelete time.Time
}

type LastMessageQueues []LastMessage

func (l *LastMessageQueues) Add(messageID int64, chatId int64, delayTime int) {
	*l = append(*l, LastMessage{
		MessageID:  &messageID,
		ChatID:     &chatId,
		TimeDelete: time.Now().Add(time.Duration(delayTime) * time.Second),
	})

}

// Size возвращает размер очереди
func (l *LastMessageQueues) Size() int {
	return len(*l)
}

func (l *LastMessageQueues) IsEmpty() bool {
	return l.Size() == 0
}

func (l *LastMessageQueues) RemoveId(id int) {
	if id+1 > l.Size() {
		return
	}
	// Вырежим укзанный ID
	//(*l)[id] = LastMessage{}
	*l = append((*l)[:id], (*l)[id+1:]...)

	return
}
