package main

import (
	"log"
	"os"

	"github.com/ark-go/fibergo/internal/db"
	"github.com/ark-go/fibergo/internal/tbot"
)

func init() {

}

func main() {
	log.Println("Бот")
	pg, err := db.StartPostgres()
	if err != nil {
		os.Exit(1)
	}
	tbot.Init(pg)
}
