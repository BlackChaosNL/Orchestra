package web

import (
	"os"
	"sync"
	"time"

	"github.com/BlackChaosNL/Orchestra/config"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/responsetime"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/kataras/golog"
)

const idleTimeout time.Duration = 5 * time.Second

var webAppPort string = config.Config("ORCHESTRA_WEB_PORT", ":9800")
var sessionStore *session.Store

var app *fiber.App

func StartWebServer(wg *sync.WaitGroup) {
	defer wg.Done()

	app = fiber.New(fiber.Config{IdleTimeout: idleTimeout})
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, //
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		Next:         nil,
	}))
	app.Use(csrf.New())
	app.Use(responsetime.New())

	app.Use(csrf.New(csrf.Config{
		CookieName:        "__Orchestra-csrf_",
		CookieSecure:      true,
		CookieHTTPOnly:    false,
		CookieSameSite:    "Lax",
		CookieSessionOnly: true,
		Extractor:         extractors.FromHeader("X-Csrf-Token"),
		Session:           sessionStore,
	}))

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

	golog.Fatal(app.Listen(webAppPort))
}

func StopWebService() {
	app.Shutdown()
}
