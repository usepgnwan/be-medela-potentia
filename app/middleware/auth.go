package middleware

import (
	"be-medela-potentia/app/helpers"
	"be-medela-potentia/app/models"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(helpers.Response{
				Success: false,
				Error:   "Invalid Token",
			})
		}

		helpers.InitAPI("", nil)

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims := &models.JwtUser{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("AUTHSECRETKEY")), nil
			},
		)

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(helpers.Response{
				Success: false,
				Error:   "Invalid Token",
			})
		}

		c.Locals("user", claims)

		return c.Next()
	}
}
