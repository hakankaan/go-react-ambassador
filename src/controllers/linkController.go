package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/gofiber/fiber/v2"
)

func Link(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var links []models.Link

	database.DB.Where("user_id = ?", id).Find(&links)

	return c.JSON(&links)
}
