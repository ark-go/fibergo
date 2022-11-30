package serverf

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func init() {

}

type ServerFiber struct {
	// собственно сам fiber
	App *fiber.App
	// адрес перенаправления запроса
	ProxyHost string
}

func (sf *ServerFiber) Start() {

	if fiber.IsChild() == true {
		log.Println("Старт  ребенком --------------->")
	} else {
		log.Println("Старт --------------->")
	}
	sf.App = fiber.New(fiber.Config{
		// Если установлено значение true, это порождает несколько процессов Go, прослушивающих один и тот же порт.
		//Prefork: true,
		// Если установлено значение true, включает маршрутизацию с учетом регистра. Например. '/FoO' и '/foo'
		CaseSensitive: false,
		// Если установлено значение true, маршрутизатор рассматривает «/foo» и «/foo/» как разные
		StrictRouting: false,
		// header - server name
		ServerHeader: "Ark Fi",
		AppName:      "Тестик v1.0.1",
	})
	// micro := fiber.New(fiber.Config{
	// 	ServerHeader: "Копия",
	// })
	// sf.App.Mount("/apix2", micro)
	// micro.Get("/apix", func(c *fiber.Ctx) error {
	// 	return c.SendString("Привет привет Micro")
	// })
	sf.ProxyHost = "http://127.0.0.1:8877"

	sf.InitRoutes()
	sf.App.Hooks().OnListen(func() error {
		log.Println("Сервер запущен")
		return nil
	})
	sf.App.Hooks().OnShutdown(func() error {
		log.Println("Shutdown server")
		return nil
	})
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	serverShutdown := make(chan struct{})
	go func() {
		_ = <-c
		log.Println("Попросим сервер заткнуться..")
		_ = sf.App.Shutdown()
		serverShutdown <- struct{}{}
	}()
	// sf.App.Listen(":8018")
	if err := sf.App.Listen(":8018"); err != nil {
		log.Panic(err)
	}
	<-serverShutdown
	log.Println("Running cleanup tasks...")
}
