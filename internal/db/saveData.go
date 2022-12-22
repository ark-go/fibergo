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

	rows, err := pg.Pool.Query(context.Background(), query, int64(usr.GetUserId()), int64(usr.GetChatId()), txt2)
	if err != nil {
		log.Fatal("Ошибка запроса DB query id:", int64(usr.GetUserId()), " err:", err.Error())
	}
	rows.Close()
	log.Println("База данных save? id:", int64(usr.GetUserId()))
}

// работает bytea  доставать - select encode(gobbyte, 'base64') from userdata
// func (pg *Pg) SaveDatabyte(usr *userdata.User) {
// 	query := `--sql
// 	INSERT INTO userdata (userid, chatid, gobdata, gobbyte) VALUES ($1, $2, $3, decode($3, 'base64'))
// 	ON CONFLICT (userid,chatid) DO UPDATE SET gobdata = EXCLUDED.gobdata, gobbyte = EXCLUDED.gobbyte, datechange = CURRENT_TIMESTAMP;
// 	`
// 	txt2 := usr.ToGOB64()

// 	rows, err := pg.Pool.Query(context.Background(), query, int64(usr.GetUserId()), int64(usr.GetChatId()), txt2)
// 	if err != nil {
// 		log.Fatal("Ошибка запроса DB query id:", int64(usr.GetUserId()), " err:", err.Error())
// 	}
// 	rows.Close()
// 	log.Println("База данных save? id:", int64(usr.GetUserId()))
// }
