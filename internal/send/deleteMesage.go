package send

import (
	"log"
	"time"

	"github.com/ark-go/fibergo/internal/msgtypes"
	"github.com/nickname76/telegrambot"
)

/*
		удаление собщений отправленных пользователю из очереди

	    если сообщения ставились в очередь при отправке
		BUG проверить когда они ставятся точно
*/
func (s *Send) DeleteMessage() error {
	log.Println("В базе ", len(s.User.UserData.MsgQueues), "записей для удаления")
	//for i, usrmes := range s.User.UserData.MsgQueues
	// будем стирать из slice поэтому задом наперед читаем
	for i := len(s.User.UserData.MsgQueues) - 1; i >= 0; i-- {
		xx := s.User.UserData.MsgQueues[i]
		if xx.TimeDelete.Before(time.Now()) {
			//log.Println("for del", *usrmes.ChatID, *usrmes.MessageID, usrmes.TimeDelete)

			s.api.DeleteMessage(&telegrambot.DeleteMessageParams{
				ChatID:    telegrambot.ChatID(*xx.ChatID),
				MessageID: telegrambot.MessageID(*xx.MessageID),
			})
			//INFO удаляем в любом случае, мы не выяняем причину но осталяя сообщение мы его больще никогда не сотрем, если оно уже стерто
			s.User.UserData.MsgQueues.RemoveId(i)

			//delete(s.User.UserData.MsgQueues, usrmes)
		}
		log.Println("Удаляем из списка", i)
	}
	//log.Println("Осталось задания для удаления:", usr.LastMsgQueues.Size())
	return nil
}

// удаление сообщения, текущего-последнего поступившего от пользователя
// кроме ответов кнопкой Inline (CallbackQuery) !
func (s *Send) DeleteMessageUser() error {
	if s.User.Info.UpdateType == msgtypes.Upd_CallbackQuery {
		return nil
	}

	log.Println("удаляем текущее, последнее входящее сообщение пользователя")
	err := s.api.DeleteMessage(&telegrambot.DeleteMessageParams{
		ChatID:    s.User.Msg.Chat.ID,
		MessageID: s.User.Msg.MessageID,
	})
	if err != nil {
		//log.Println("Ошибка удаления сообщения ", msg.MessageID, err.Error())
	}
	return nil
}

// удаление сообщений с Меню-кнопками в зависимости от разного источника где был юзер чат бот и тд
func (s *Send) deleteMessageMenu() error {
	log.Println("на удаление меню", len(s.User.UserData.MenuAll), "записей")
	for _, usrmes := range s.User.UserData.MenuAll {
		if usrmes.ChatID == s.User.GetChatId() {
			err := s.api.DeleteMessage(&telegrambot.DeleteMessageParams{
				ChatID:    telegrambot.ChatID(usrmes.ChatID),
				MessageID: telegrambot.MessageID(usrmes.MessageID),
			})
			if err == nil {
				delete(s.User.UserData.MenuAll, usrmes.ChatID)
			}
		}
	}
	return nil
}
