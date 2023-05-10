package routes

import (
	"github.com/NimishKashyap/tryout-go/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	app.Get("/main/:id", controllers.MainController)
	app.Post("/main", controllers.BodyController)
}
