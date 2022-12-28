package db

import (
	"context"
	"errors"
	"log"

	//"github.com/ark-go/fibergo/internal/tbot"
	"github.com/ark-go/fibergo/internal/userdata"
	"github.com/nickname76/telegrambot"
)

func (pg *Pg) LoadData(msg *telegrambot.Message) (*userdata.User, error) {
	//!! создаем новый  Обязательно
	user := userdata.InitUser(msg)
	if msg.From == nil {
		return nil, errors.New("нет From у Msd")
	}
	query := `--sql
	SELECT gobdata FROM userdata 
	WHERE userid = $1 AND chatid = $2
	`

	rows, err := pg.Pool.Query(context.Background(), query, int64(user.GetUserId()), int64(user.GetChatId()))
	if err != nil {
		log.Fatal("Ошибка запроса DB LoadData  id:", int64(user.GetUserId()), " err:", err.Error())
	}
	defer rows.Close()
	if rows.Err() != nil {
		log.Println("Err DB rows:", rows.Err())
		return user, nil
	}
	var res string = ""
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println("LoadData: ошибка при обходе набора данных")
			return user, nil
		}
		res = values[0].(string)
	}

	if res == "" {
		// Не найдено
		log.Printf("Новый клиент %+v", user.GetUserId())
		return user, nil
	}

	// и записываем в него данные из базы
	user.FromGOB64(res, msg)

	log.Printf("База данных LoadData : %d, chat %d", user.GetUserId(), user.GetChatId())

	return user, nil
}
