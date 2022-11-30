package tbot

import (
	"log"
	"time"

	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

// удаление собщений отправленных пользователю из очереди
func (b *Bot) deleteMessage(usr *userdata.User) error {
	for i, usrmes := range usr.Last.MsgQueues {
		if usrmes.TimeDelete.Before(time.Now()) {
			//log.Println("for del", *usrmes.ChatID, *usrmes.MessageID, usrmes.TimeDelete)

			b.Api.DeleteMessage(&telegrambot.DeleteMessageParams{
				ChatID:    telegrambot.ChatID(*usrmes.ChatID),
				MessageID: telegrambot.MessageID(*usrmes.MessageID),
			})
			//INFO удаляем в любом случае, мы не выяняем причину но осталяя сообщение мы его больще никогда не сотрем, если оно уже стерто
			usr.Last.MsgQueues.RemoveId(i)
		}
	}
	//log.Println("Осталось задания для удаления:", usr.LastMsgQueues.Size())
	return nil
}

// удаление сообщение из параметра
func (b *Bot) deleteMessageUser(msg *telegrambot.Message) error {
	err := b.Api.DeleteMessage(&telegrambot.DeleteMessageParams{
		ChatID:    msg.Chat.ID,
		MessageID: msg.MessageID,
	})
	if err != nil {
		//log.Println("Ошибка удаления сообщения ", msg.MessageID, err.Error())
	}
	return nil
}

// удаление сообщений с Меню-кнопками в зависимости от разного источника где был юзер чат бот и тд
func (b *Bot) deleteMessageMenu(usr *userdata.User) error {
	for _, usrmes := range usr.Last.MenuAll {
		log.Println("меню было", len(usr.Last.MenuAll))
		if usrmes.ChatID == usr.MsgChatId() {
			err := b.Api.DeleteMessage(&telegrambot.DeleteMessageParams{
				ChatID:    telegrambot.ChatID(usrmes.ChatID),
				MessageID: telegrambot.MessageID(usrmes.MessageID),
			})
			if err == nil {
				delete(usr.Last.MenuAll, usrmes.ChatID)
			}
		}
		log.Println("меню стало", len(usr.Last.MenuAll))
	}
	return nil
}
