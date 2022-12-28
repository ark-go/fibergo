package send

import (
	"log"
	"net/url"
	"os"

	"github.com/nickname76/telegrambot"
)

func (s *Send) PhotoGroup(knopki *telegrambot.InlineKeyboardMarkup, filePath [][]string) error {
	var mediaArr []*telegrambot.InputMedia
	var y = 0
	//var mediaArrFull = map[int][]*telegrambot.InputMedia{}
	for i, val := range filePath {
		if i%2 == 0 {
			continue
		}
		if y++; y > 9 {
			break
		}
		mediaArr = append(mediaArr, &telegrambot.InputMedia{
			Type:      telegrambot.InputMediaTypePhoto,
			Media:     telegrambot.FileURL(val[0]),
			Caption:   val[1],
			ParseMode: telegrambot.ParseModeHTML,
		})
		//log.Println("photo:", val)
	}
	if len(mediaArr) > 0 {
		mess, err := s.api.SendMediaGroup(&telegrambot.SendMediaGroupParams{
			ChatID: s.User.GetChatId(),
			Media:  mediaArr,
			//ProtectContent: true,
		})
		if err != nil {
			log.Println("SendMediaGroup Error:", err)
			return err
		} else {
			for _, v := range mess {
				s.User.UserData.MsgQueues.Add(int64(v.MessageID), int64(s.User.GetChatId()), 1)
			}
		}
	} else {
		log.Println("картинок нет для групповухи")
	}
	return nil
}
func (s *Send) Photo(knopki *telegrambot.InlineKeyboardMarkup, filePath, fileId string, caption ...string) (*telegrambot.Message, error) {
	if len(caption) == 0 {
		//caption = append(caption, "<b>Я тут бот.</b>\n")
		caption = append(caption, "")
	}
	// ---------------------- Картинка по URL  -------------------------
	u, er := url.Parse(filePath) // request хочет схему, короче проверять на схему
	if er == nil {
		// если это Url т.е. есть схема http... но какая схема не проверяем
		log.Println("это URL? >", u.IsAbs())
		if u.IsAbs() {
			// Если задан filePath = URL пробуем отправить с ним
			params := &telegrambot.SendPhotoParams{
				ChatID: s.User.GetChatId(),
				Photo:  telegrambot.FileURL(filePath),
				//	ProtectContent: true,
				ParseMode: telegrambot.ParseModeHTML,
				Caption:   caption[0],
			}
			if knopki != nil {
				params.ReplyMarkup = knopki
			}
			msg, err := s.api.SendPhoto(params)
			if err != nil {
				return nil, err
			} else {
				s.User.UserData.MsgQueues.Add(int64(msg.MessageID), int64(s.User.GetChatId()), 1)
				log.Println("Отправили картинку по URL", len(s.User.UserData.MsgQueues))
				return msg, nil
			}
		}
	}
	//------------------- Картинка по ID ---------------
	// Если задан fileID пробуем отправить с ним
	if fileId != "" {
		msg, err := s.api.SendPhoto(&telegrambot.SendPhotoParams{

			ChatID:         s.User.GetChatId(),
			Photo:          telegrambot.FileID(fileId),
			ProtectContent: true,
			ReplyMarkup:    knopki,
			ParseMode:      telegrambot.ParseModeHTML,
			Caption:        caption[0],
		})
		if err == nil {
			log.Println("Отправили картинку по ID")
			return msg, nil
		}
	}
	// ---------------------- Картинка из файла ------------------
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	photo := &telegrambot.FileReader{
		Name:   "test",
		Reader: file,
	}
	msg, err := s.api.SendPhoto(&telegrambot.SendPhotoParams{

		ChatID:         s.User.GetChatId(),
		Photo:          photo,
		ProtectContent: true,
		ReplyMarkup:    knopki,
		ParseMode:      telegrambot.ParseModeHTML,
		Caption:        caption[0],
	})

	if err == nil {
		id := msg.Photo[0].FileID // этот ИД можно отправить вместо картинки?
		log.Println("Картинка отправлена ID:", id)
	}
	return msg, err
}
