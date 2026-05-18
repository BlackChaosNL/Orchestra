package middleware

import (
	"fmt"
	"strings"

	"github.com/BlackChaosNL/Orchestra/config"
	"github.com/gofiber/fiber/v3"

	jwtware "github.com/gofiber/contrib/v3/jwt"
)

func Protected() fiber.Handler {

}

func OAUTHProtected() fiber.Handler {

}

func JWTProtected() fiber.Handler {
	secret := config.GetStrFromEnv("ORCHESTRA_API_SECRET_KEY", "")

	if secret == "" {
		fmt.Errorf("ORCHESTRA_API_SECRET_KEY has not been properly filled.")
	}

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(c fiber.Ctx, err error) error {
			status := fiber.StatusUnauthorized
			message := "Invalid or expired JWT"

			if strings.Contains(strings.ToLower(err.Error()), "missing or malformed jwt") {
				status = fiber.StatusBadRequest
				message = "Missing or malformed JWT"
			}

			return c.Status(status).JSON(fiber.Map{
				"status":  "error",
				"message": message,
				"data":    nil,
			})
		},
	})

}
