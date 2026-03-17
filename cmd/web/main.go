package web

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/kataras/golog"
)

var app *fiber.App

const idleTimeout = 5 * time.Second

func StartWebServer(wg *sync.WaitGroup) {
	defer wg.Done()

	app = fiber.New(fiber.Config{IdleTimeout: idleTimeout})
	app.Use(golog.New())
	app.Use(cors.New())

	app.Get("/", static.New("./ui/dist"), static.Config{
		MaxAge: 0,
	})

	app.Get("/static/*", static.New("./ui/dist/static", static.Config{
		MaxAge: 31536000, // 1 year
	}))

	golog.Fatal(app.Listen(":9800"))
}

func StopAPIService() {
	app.Shutdown()
}
