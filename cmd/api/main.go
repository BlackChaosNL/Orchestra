package api

import (
	"sync"
	"time"

	"github.com/BlackChaosNL/Orchestra/cmd/api/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/kataras/golog"
)

var app *fiber.App

const idleTimeout = 5 * time.Second

func StartAPIServer(wg *sync.WaitGroup) {
	defer wg.Done()
	app = fiber.New(fiber.Config{IdleTimeout: idleTimeout})
	app.Use(golog.New())
	app.Use(cors.New())

	g := app.Group("/api/v1/")

	routes.GlobalRouter(g)

	golog.Fatal(app.Listen(":9810"))
}

func StopAPIService() {
	app.Shutdown()
	golog.Print("Test")
}
