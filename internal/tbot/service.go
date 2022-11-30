package tbot

import (
	"log"
	"os"
	"os/signal"

	"github.com/ark-go/fibergo/internal/db"
	"github.com/nickname76/telegrambot"
)

type Bot struct {
	Api *telegrambot.API
	Me  *telegrambot.User
	Pg  *db.Pg
}

func Init(pg *db.Pg) {

	api, me, err := telegrambot.NewAPI("5816387767:AAG-2KPVIppM1PemOzwa6RaGMtSXOODjppM")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	bot := &Bot{
		Api: api,
		Me:  me,
		Pg:  pg,
	}

	stop := telegrambot.StartReceivingUpdates(api, func(update *telegrambot.Update, err error) {
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}
		bot.Update(update)
	})

	log.Printf("Старт бота: %v", bot.Me.Username)

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt)

	<-exitCh

	// Waits for all updates handling to complete
	stop()
}
