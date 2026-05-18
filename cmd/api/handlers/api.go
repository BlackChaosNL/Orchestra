package handlers

import "github.com/gofiber/fiber/v3"

func Root(c fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"ping": "pong!",
	})
}

func Version(c fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"version": "0.0.0",
	})
}
