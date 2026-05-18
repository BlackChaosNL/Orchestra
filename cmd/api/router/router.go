package router

import (
	"github.com/BlackChaosNL/Orchestra/cmd/api/handlers"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Router struct {
	app *fiber.App
	db  *gorm.DB
}

func CreateRouter(app *fiber.App, db *gorm.DB) *Router {
	return &Router{
		app: app,
		db:  db,
	}
}

func (r *Router) SetupRoutes() {
	// Default routes are under `${url}/api/v1`
	v1 := r.app.Group("/v1")
	v1.Get("/", handlers.Root)
	v1.Get("/version", handlers.Version)

}
