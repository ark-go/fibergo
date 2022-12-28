package send

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ark-go/fibergo/internal/utils"
	"github.com/nickname76/telegrambot"
)

/*
закачка файла
*/
func (s *Send) LoadFileFromFileID(fileId telegrambot.FileID) (string, error) {

	filePhoto, _ := s.api.GetFile(&telegrambot.GetFileParams{
		FileID: fileId, // s.User.Msg.Photo[len(s.User.Msg.Photo)-1].FileID,
	})

	fileName := filepath.Base(filePhoto.FilePath)
	ext := filepath.Ext(filePhoto.FilePath)
	fmt.Println("Файл", filePhoto.FilePath, "Ext", ext)
	fileUrl := "https://api.telegram.org/file/bot" + os.Getenv("TG_Bot") + "/" + filePhoto.FilePath

	path := filepath.Join(utils.ExecDir, fileName)
	endDown := make(chan bool)
	notify := make(chan *utils.WriteCounter)
	go utils.DownloadFile(path, fileUrl, endDown, notify)
exit:
	for {
		select {
		case <-endDown:
			// переводим строку
			fmt.Printf("\n")
			break exit
		case res := <-notify:
			// печатаем без перевода строки
			space := strings.Repeat(" ", 35) // для стирания с терминала старой строки
			// Вернуться снова и вывести текущий статус загрузки
			fmt.Printf("\r%s\rЗакачали: %d%% - %s (%s) готово", space, res.TotalPrc, utils.ByteCountSI(res.Total), utils.ByteCountIEC(res.Total))
		}
	}
	return path, nil
}
