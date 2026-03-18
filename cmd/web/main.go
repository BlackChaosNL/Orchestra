package web

import (
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/responsetime"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/kataras/golog"
)

var app *fiber.App

const idleTimeout = 5 * time.Second

func StartWebServer(wg *sync.WaitGroup) {
	defer wg.Done()

	app = fiber.New(fiber.Config{IdleTimeout: idleTimeout})
	app.Use(cors.New())
	app.Use(responsetime.New())

	app.Get("/assets*", static.New("", static.Config{
		MaxAge:   31536000, // 1 year
		FS:       os.DirFS("cmd/web/ui/dist/assets"),
		Browse:   true,
		Compress: true,
	}))

	app.Get("/*", func(c fiber.Ctx) error {
		return c.SendFile("cmd/web/ui/dist/index.html", fiber.SendFile{
			MaxAge:   0, // Force fresh index.
			Compress: true,
		})
	})

	golog.Fatal(app.Listen(":9800"))
}

func StopAPIService() {
	app.Shutdown()
}
