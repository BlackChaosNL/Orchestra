package api

import (
	"os"
	"sync"
	"time"

	"github.com/BlackChaosNL/Orchestra/cmd/api/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/responsetime"
	"github.com/kataras/golog"
)

const idleTimeout time.Duration = 5 * time.Second
const defaultPort string = ":9810"

var userPort string = os.Getenv("ORCHESTRA_API_PORT")
var app *fiber.App

func StartAPIServer(wg *sync.WaitGroup) {
	defer wg.Done()

	app = fiber.New(fiber.Config{IdleTimeout: idleTimeout})
	app.Use(cors.New())
	app.Use(responsetime.New())

	// Set global group to /api, let routes choose version.
	prefix := app.Group("/api")

	routes.GlobalRouter(prefix)

	golog.Fatal(app.Listen(func() string {
		if userPort == "" {
			return defaultPort
		} else {
			return userPort
		}
	}()))
}

func StopAPIService() {
	app.Shutdown()
}
