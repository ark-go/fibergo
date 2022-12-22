package send

import (
	"errors"
	"log"
	"os"

	"github.com/nickname76/telegrambot"
)

/*
	    Отпрака группы фоток
		нужен массив путей к файлам
*/
func (s *Send) SendGroupPhoto(filePath []string) ([]*telegrambot.Message, error) {
	filePointer := []*os.File{}

	for i := 0; i < len(filePath); i++ {
		file, err := os.Open(filePath[i])
		if err != nil {
			log.Println("SendGroupPhoto: Ошибка открытия файла: ", err)
			continue
		}
		// !! пример, оставить одну строку close
		defer func(filed *os.File, m int) {
			err := filed.Close()
			log.Println("Закрыли файл", m, err)
		}(file, i) // необходимо закрепить переменные в функции, иначе возьмет из последнего for цикла
		filePointer = append(filePointer, file)
	}
	// не открыли ни одного файла
	if len(filePointer) < 1 {
		return nil, errors.New("SendGroupPhoto: Не чего отправить.")
	}

	// собираем группу фоток
	MediaGroup := []*telegrambot.InputMedia{}
	for _, file := range filePointer {
		MediaGroup = append(MediaGroup, &telegrambot.InputMedia{
			Type: telegrambot.InputMediaTypePhoto,
			Media: &telegrambot.FileReader{
				Name:   "test",
				Reader: file,
			},
		},
		)
	}
	// отправляем группу
	msgout, err := s.api.SendMediaGroup(&telegrambot.SendMediaGroupParams{
		ChatID:         s.User.GetChatId(),
		Media:          MediaGroup,
		ProtectContent: true,
	})
	if err != nil {
		return nil, err
	}
	return msgout, nil
}
