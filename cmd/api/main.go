package api

import (
	"sync"
	"time"

	"github.com/BlackChaosNL/Orchestra/cmd/api/internal"
	"github.com/BlackChaosNL/Orchestra/cmd/api/router"
	"github.com/BlackChaosNL/Orchestra/config"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/responsetime"
)

const idleTimeout time.Duration = 5 * time.Second

var apiAppPort string = config.GetStrFromEnv("ORCHESTRA_API_PORT", ":9810")
var tofu, tofuDir, workPath = internal.GetTofu(config.GetStrFromEnv("ORCHESTRA_API_OPENTOFU_VERSION", "1.11.5"))
var app *fiber.App

func StartAPIServer(wg *sync.WaitGroup) {
	defer wg.Done()

	db := internal.SetupDB()
	internal.LoadTables(db)

	app = fiber.New(fiber.Config{IdleTimeout: idleTimeout})
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, //
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		Next:         nil,
	}))
	app.Use(recover.New())
	app.Use(responsetime.New())

	// Set global group to /api, let routes choose version.
	app.Group("/api")
	r := router.CreateRouter(app, db)
	r.SetupRoutes()

	app.Listen(apiAppPort)
}

func StopAPIService() {
	app.Shutdown()
	internal.RemoveFolder(tofuDir)
	internal.RemoveFolder(workPath)
}
