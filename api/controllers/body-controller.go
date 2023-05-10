package controllers

import (
	"log"

	"github.com/NimishKashyap/tryout-go/api/db"
	"github.com/NimishKashyap/tryout-go/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func BodyController(c *fiber.Ctx) error {
	p := new(models.Users)
	dbConn := db.GetDB()

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Age)
	log.Println(p.Name)

	result := dbConn.Create(&p)

	log.Println(p.ID)
	log.Println(result.Error)
	log.Println(result.RowsAffected)

	return c.Status(201).SendString("CREATED")
}
