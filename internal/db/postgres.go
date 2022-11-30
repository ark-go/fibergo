package db

import (
	"context"
	"log"
	"os"

	"github.com/ark-go/fibergo/internal/utils"
	"github.com/fatih/color"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pg struct {
	Pool *pgxpool.Pool
}

func StartPostgres() (*Pg, error) {
	pg := &Pg{}
	config, err := pgxpool.ParseConfig(os.Getenv("PG_DatabaseStr"))
	if err != nil {
		log.Println("Ошибка парсинга конфига pgxpool")
		return nil, err
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		log.Println("Коннект из пула : ", config.MinConns, config.MaxConns)
		return nil
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Println("Не удалось создать пул соединений: ", err.Error())
		//os.Exit(1)
		return nil, err

	}
	pg.Pool = dbpool

	sizeDB := pg.sizeDB(dbpool.Config().ConnConfig.Database)
	color.Set(color.FgGreen)
	log.Println("Pg подключились, DB Name:", dbpool.Config().ConnConfig.Config.User, utils.ByteCountSI(sizeDB))
	color.Unset()
	return pg, nil
}

// размер базы
func (pg *Pg) sizeDB(nameDB string) int64 {
	var sz int64
	err := pg.Pool.QueryRow(context.Background(), "select "+"pg_database_size('"+nameDB+"')").Scan(&sz)
	if err != nil {
		log.Println("QueryRow failed: ", err)
		return 0
	}
	return sz
}

// закроем подключение к Pg
func (pg *Pg) Close() {
	pg.Pool.Close()
}

/*
su postgres -c psql
create database dbname with encoding='UNICODE';
create user dbuser with password 'dbpass';
grant all privileges on database dbname to dbuser;

*/
