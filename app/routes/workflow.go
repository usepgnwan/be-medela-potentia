package routes

import (
	"be-medela-potentia/app/controllers"
	"be-medela-potentia/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WorkflowRoute(r fiber.Router) {
	app := r.Group("workflows")
	app.Get("/", controllers.GetWorkflow)
	app.Get("/:id", controllers.GetDetailWorkflow)
	app.Post("/", middlewares.UserAuthorization(), controllers.PostWorkflow)
}
