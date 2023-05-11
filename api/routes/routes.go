package routes

import (
	"github.com/NimishKashyap/tryout-go/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	app.Get("/main/:cid", controllers.MainController)
	app.Post("/main", controllers.BodyControllerTx)
	app.Post("/company", controllers.CreateCompanyController)
}
