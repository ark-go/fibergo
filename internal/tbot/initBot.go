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
	Me  *telegrambot.User // сам бот
	Pg  *db.Pg
	// program func()
}

var Api *telegrambot.API

func InitBot(pg *db.Pg) {
	api, me, err := telegrambot.NewAPI(os.Getenv("TG_Bot"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	Api = api
	var bot = &Bot{
		Api: api,
		Me:  me,
		Pg:  pg, //TODO Убрать
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

// func StartBot(bot *Bot, pg *db.Pg) {
// 	stop := telegrambot.StartReceivingUpdates(bot.Api, func(update *telegrambot.Update, err error) {
// 		if err != nil {
// 			log.Printf("Error: %v", err)
// 			return
// 		}
// 		bot.Update(update)
// 	})

// 	log.Printf("Старт бота: %v", bot.Me.Username)

// 	exitCh := make(chan os.Signal, 1)
// 	signal.Notify(exitCh, os.Interrupt)

// 	<-exitCh

// 	// Waits for all updates handling to complete
// 	stop()
// }
