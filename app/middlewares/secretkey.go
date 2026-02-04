package middlewares

import (
	"be-medela-potentia/app/helpers"
	"os"

	"github.com/gofiber/fiber/v2"
)

func ApiKeyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("x-api-key")

		if apiKey == "" || apiKey != os.Getenv("APIKEY") {
			return c.Status(fiber.StatusUnauthorized).JSON(helpers.Response{
				Success: false,
				Error:   "Invalid or missing API key",
			})
		}
		return c.Next()
	}
}
