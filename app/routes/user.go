package routes

import (
	"be-medela-potentia/app/controllers"
	"be-medela-potentia/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router) {
	app := r.Group("users")
	app.Get("/", controllers.GetUser)
	app.Get("/check-jwt", middleware.UserAuthorization(), controllers.ClaimJwt)
	app.Post("/", controllers.PostUser)
	app.Post("/login", controllers.UserLogin)
}
