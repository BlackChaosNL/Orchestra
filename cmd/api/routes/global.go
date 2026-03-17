package routes

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func GlobalRouter(r fiber.Router) {
	r.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"ping": "pong!",
		})
	})

	r.Get("/time", func(c fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"current_time": time.Now(),
		})
	})
}
