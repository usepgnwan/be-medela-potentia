package routes

import (
	"be-medela-potentia/app/controllers"
	"be-medela-potentia/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RequestRoute(r fiber.Router) {
	app := r.Group("request")
	app.Get("/:id", controllers.GetDetailRequest)
	app.Post("/", middlewares.UserAuthorization(), controllers.PostRequest)
	app.Post("/:id/reject", middlewares.UserAuthorization(), controllers.RejectRequest)
	app.Post("/:id/approve", middlewares.UserAuthorization(), controllers.ApproveRequest)

}
