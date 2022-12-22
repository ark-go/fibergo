package userdata

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"

	"github.com/nickname76/telegrambot"
)

func (u *User) ToGOB64() string {
	// упаковываем и храним только u.UserData
	b := &bytes.Buffer{}
	//	gob.Register(u.LastMsgQueues.GetValues())
	// gob.Register(u)
	// gob.Register(u.LastMsgQueues)
	e := gob.NewEncoder(b)

	err := e.Encode(u.UserData)
	if err != nil {
		fmt.Println(`Ошибка failed gob Encode`, err, err.Error())
	}
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func (u *User) FromGOB64(str string, msg *telegrambot.Message) {
	// распаковываем и добавляем только u.UserData
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(`не раскодировать данные из базы`, err)
		u = InitUser(msg)
	}
	b := bytes.Buffer{}
	b.Write(by)
	// gob.Register(u)
	// gob.Register(u.LastMsgQueues)

	d := gob.NewDecoder(&b)
	err = d.Decode(&u.UserData)
	if err != nil {
		fmt.Println(`Не раскодировать данные GOB`, err)
		u = InitUser(msg)
	}
	u.Olden = true
	//fmt.Println("inl257:", *u.UserData.InlineMenuAll[u.GetChatId()].MessageID)
	//u.ChangeMessage(msg) // заменяем новым сообщением
	return
}
