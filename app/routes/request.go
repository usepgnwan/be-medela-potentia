package routes

import (
	"be-medela-potentia/app/controllers"
	"be-medela-potentia/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func RequestRoute(r fiber.Router) {
	app := r.Group("request")
	app.Get("/:id", controllers.GetDetailRequest)
	app.Post("/", middleware.UserAuthorization(), controllers.PostRequest)
	app.Post("/:id/reject", middleware.UserAuthorization(), controllers.RejectRequest)
	app.Post("/:id/approve", middleware.UserAuthorization(), controllers.ApproveRequest)

}
