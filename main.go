package main

import (
	"be-medela-potentia/app/middlewares"
	"be-medela-potentia/app/routes"
	connection "be-medela-potentia/conection"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	connection.ConnectDB()
	app := fiber.New()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3011"
	}
	app.Use(middlewares.Logger())
	routes.InitRoutes(app)
	app.Listen(fmt.Sprintf(":%s", port))
}
