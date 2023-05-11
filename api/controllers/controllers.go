package controllers

import (
	"log"

	"github.com/NimishKashyap/tryout-go/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func MainController(c *fiber.Ctx) error {
	id, err := c.ParamsInt("cid")
	if err != nil {
		log.Fatalf("Could not find param")
	}

	emp := models.Users{Name: "Nimish", Age: 10, CompanyID: id}

	return c.JSON(emp)
}
