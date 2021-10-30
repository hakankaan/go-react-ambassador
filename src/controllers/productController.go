package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

func CreateProducts(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	id, _ := c.ParamsInt("id")

	product.Id = uint(id)

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	id, _ := c.ParamsInt("id")

	product.Id = uint(id)

	database.DB.Model(&product).Updates(&product)

	return c.JSON(&product)
}

func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product

	id, _ := c.ParamsInt("id")

	product.Id = uint(id)

	database.DB.Delete(&product)

	return nil
}
