package main

import (
	"be-medela-potentia/app/middlewares"
	"be-medela-potentia/app/routes"
	connection "be-medela-potentia/conection"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

// @title pt medela potentia tbk. API Documentation
// @version		1.0
// @description  Dokumentasi Api workflows sederhana

// @BasePath		/
// @contact.name workflows sederhana
// @contact.url http://usepgnwan.my.id
// @contact.email usepgnwan76@gmail.com

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
