package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func Cors() fiber.Handler {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	allowOrigins := os.Getenv("ALLOW_ORIGINS")
	if allowOrigins == "" {
		log.Fatalf("ALLOW_ORIGINS is not set in .env file")
	}

	corsConfig := cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		// AllowHeaders: "Origin,Content-Type,Accept,Authorization,X-Requested-With,X-API-KEY",
	}

	return cors.New(corsConfig)
}
