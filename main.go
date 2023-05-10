package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/NimishKashyap/tryout-go/api/migrations"
	"github.com/NimishKashyap/tryout-go/api/routes"
	"github.com/NimishKashyap/tryout-go/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	utils.ImportEnv()

	app := fiber.New(fiber.Config{})

	app.Get("/", healthCheck)
	app.Get("/health", healthCheck)

	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.HasPrefix(c.Path(), "api")
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowHeaders: "*",
	}))

	routes.MountRoutes(app)

	port := utils.GetPort()

	if viper.GetBool("MIGRATE") {
		migrations.Migrate()
	}

	fmt.Printf("MIGRATE: %v", viper.GetBool("MIGRATE"))

	log.Println("Listening server")
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
