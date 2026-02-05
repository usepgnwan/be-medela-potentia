package routes

import (
	"be-medela-potentia/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func WorkflowStepRoute(r fiber.Router) {
	app := r.Group("workflows")
	app.Get("/:id/step", controllers.GetDetailWorkflowStep)
	app.Post("/:id/step", controllers.PostWorkflowStep)
}
