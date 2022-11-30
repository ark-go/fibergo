package serverf

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/valyala/fasthttp"
)

func (sf *ServerFiber) InitRoutes() {
	sf.App.Use(func(c *fiber.Ctx) error {
		// Set a custom header on all responses: только Post если в Proxy?
		// вероятно, что лучше выставить можно в роуте. не понятно почему тут при Get не ставится в Proxy
		c.Set("content-type", "text/plain; charset=utf-8") // гдето оно по умолчанию ставится
		c.Set("info2", "karamba")
		c.Response().Header.Del(fiber.HeaderServer) // удалит Header "server"
		return c.Next()
	})
	proxy.WithClient(&fasthttp.Client{
		NoDefaultUserAgentHeader: true, // true исключает заголовок User-Agent
		DisablePathNormalizing:   true, // отключает нормализацию путей, отправляет как есть.
	})
	sf.App.Get("/metrics", monitor.New(monitor.Config{Title: "Мой серверок"}))

	sf.App.Post("/api/*", sf.ProxyFunc(true))

	sf.App.Get("/apiv", func(c *fiber.Ctx) error {
		return c.SendString("Привет привет Privet")
	})

	socket := sf.App.Group("/socket.io/*")
	socket.Get("*", sf.ProxyFunc(true))
	socket.Post("*", sf.ProxyFunc(true))

	sf.App.Get("/*", sf.ProxyFunc(true))

}
