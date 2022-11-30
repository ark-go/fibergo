package db

import (
	"context"
	"log"

	"github.com/ark-go/fibergo/internal/userdata"
)

func (pg *Pg) SaveData(usr *userdata.User) {
	query := `--sql
	INSERT INTO userdata (userid, chatid, gobdata) VALUES ($1, $2, $3)
	ON CONFLICT (userid,chatid) DO UPDATE SET gobdata = EXCLUDED.gobdata, datechange = CURRENT_TIMESTAMP;
	`
	txt2 := usr.ToGOB64()

	rows, err := pg.Pool.Query(context.Background(), query, int64(usr.UserId), int64(usr.ChatId), txt2)
	if err != nil {
		log.Fatal("Ошибка запроса DB query id:", int64(usr.UserId), " err:", err.Error())
	}
	rows.Close()
	log.Println("База данных save? id:", int64(usr.UserId))
}
