package controllers

import (
	"log"

	"github.com/NimishKashyap/tryout-go/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func MainController(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatalf("Could not find param")
	}

	log.Println(id)

	emp := models.Users{Name: "Nimish", Age: 10}

	return c.JSON(emp)
}
