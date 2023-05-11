package controllers

import (
	"log"

	"github.com/NimishKashyap/tryout-go/api/db"
	"github.com/NimishKashyap/tryout-go/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func BodyControllerTx(c *fiber.Ctx) error {
	p := new(models.Users)

	if err := c.BodyParser(&p); err != nil {
		log.Fatal("Could not parse the body")
	}

	dbConn := db.GetDB()

	tx := dbConn.Begin()

	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		log.Fatalf("Error during creation: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Fatalf("Error during commit proces: %v", err)
	}

	return c.Status(201).SendString("Created through transactions")
}
