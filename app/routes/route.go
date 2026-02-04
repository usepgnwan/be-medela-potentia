package routes

import (
	"be-medela-potentia/app/middlewares"
	_ "be-medela-potentia/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func InitRoutes(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome, on api page approval")
	})
	api := r.Group("/api")
	swaggerDoc := api.Group("/documentation")
	swaggerDoc.Get("/swagger/*", middlewares.SwagMiddleware(), swagger.HandlerDefault)
	// call route
	apiprivate := api.Use(middlewares.ApiKeyMiddleware())
	UserRoute(apiprivate)
	RoleRoute(apiprivate)
	WorkflowRoute(apiprivate)
	WorkflowStepRoute(apiprivate)
}
