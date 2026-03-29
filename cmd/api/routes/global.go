package routes

import (
	"github.com/gofiber/fiber/v3"
)

func GlobalRouter(r fiber.Router) {
	f := r.Group("/v1")

	f.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"ping": "pong!",
		})
	})

}
