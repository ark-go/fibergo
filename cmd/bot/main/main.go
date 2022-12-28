package main

import (
	//"io"
	"log"
	"os"

	//	"os"

	"github.com/ark-go/fibergo/internal/db"
	"github.com/ark-go/fibergo/internal/webs"

	_ "github.com/ark-go/fibergo/internal/programs/pDefault"
	"github.com/ark-go/fibergo/internal/tbot"
)

func init() {

}

func main() {
	// Отключает вывод лог по всему проекту
	//log.SetOutput(io.Discard)
	log.Println("Бот")
	pg, err := db.StartPostgres()
	if err != nil {
		os.Exit(1)
	}
	go webs.Start()
	tbot.InitBot(pg)
	// sv := &services.Service{}
	// sv.InitBot()
}
