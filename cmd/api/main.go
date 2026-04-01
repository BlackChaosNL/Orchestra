package api

import (
	"context"
	"sync"
	"time"

	"github.com/BlackChaosNL/Orchestra/cmd/api/routes"
	"github.com/BlackChaosNL/Orchestra/config"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/responsetime"
	"github.com/kataras/golog"
	"github.com/opentofu/tofu-exec/tfexec"
)

const idleTimeout time.Duration = 5 * time.Second

var apiAppPort string = config.Config("ORCHESTRA_API_PORT", ":9810")
var tofu, tofuDir, workPath = config.GetTofu(config.Config("ORCHESTRA_API_OPENTOFU_VERSION", "1.11.5"))
var app *fiber.App

func StartAPIServer(wg *sync.WaitGroup) {
	defer wg.Done()

	err := tofu.Init(context.Background(), tfexec.Upgrade(true))

	if err != nil {
		golog.Infof("Tofudir: %s", tofuDir)
		golog.Fatalf("Error with OpenTofu: %s", err)
	}

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
	prefix := app.Group("/api")

	routes.GlobalRouter(prefix)

	app.Listen(apiAppPort)
}

func StopAPIService() {
	app.Shutdown()
	config.RemoveFolder(tofuDir)
	config.RemoveFolder(workPath)
}
