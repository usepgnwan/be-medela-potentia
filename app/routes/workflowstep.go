package routes

import (
	"be-medela-potentia/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func WorkflowStepRoute(r fiber.Router) {
	app := r.Group("workflows-step")
	app.Get("/:id", controllers.GetDetailWorkflowStep)
	app.Post("/", controllers.PostWorkflowStep)
}
