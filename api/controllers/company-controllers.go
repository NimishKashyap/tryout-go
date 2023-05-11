package controllers

import (
	"log"

	"github.com/NimishKashyap/tryout-go/api/db"
	"github.com/NimishKashyap/tryout-go/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCompanyController(c *fiber.Ctx) error {
	var company models.Company
	dbConn := db.GetDB()
	if err := c.BodyParser(&company); err != nil {
		log.Fatal("Could not parse body")
	}
	result := dbConn.Create(&company)

	log.Println("INSERTED ID: ", company.ID)
	log.Println("INSERTED Rows: ", result.RowsAffected)

	return c.Status(201).SendString("201 Created")
}
