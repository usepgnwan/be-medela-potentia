package routes

import (
	"be-medela-potentia/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func RoleRoute(r fiber.Router) {
	app := r.Group("roles")
	app.Get("/", controllers.GetRole)
	app.Post("/", controllers.PostRole)
}
