package api

import (
	"sync"

	"github.com/gofiber/fiber/v3"
)

func StartAPIServer(wg *sync.WaitGroup) {
	defer wg.Done()
	r := fiber.New()

	r.Group("/api/v1/")

	r.Get("/api/v1/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"ping": "pong!",
		})
	})

	r.Listen(":8080")
}
