package controllers

import (
	"dynamic-pricing-api/backend-go/database"
	"dynamic-pricing-api/backend-go/models"

	"github.com/gofiber/fiber/v2"
)

// Get all products
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Find(&products)
	return c.Status(200).JSON(products)
}

// Create a new product
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Request"})
	}

	product.CalculatePrice()
	database.DB.Create(&product)

	return c.Status(201).JSON(product)
}

// Update product demand and recalculate price
func UpdateProductPrice(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Request"})
	}

	product.CalculatePrice()
	database.DB.Save(&product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	database.DB.Delete(&product)

	return c.Status(200).JSON(product)
}
