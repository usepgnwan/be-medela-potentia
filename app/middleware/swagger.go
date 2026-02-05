package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func SwagMiddleware() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Authorizer: func(username, password string) bool {
			return username == os.Getenv("USER_SWAG") &&
				password == os.Getenv("PASS_SWAG")
		},
	})
}
