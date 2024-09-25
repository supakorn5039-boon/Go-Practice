package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Product represents a product in the database.
type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// GetProducts retrieves all products from the database.
func GetProducts(c *fiber.Ctx) error {
	var products []Product
	if err := DB.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve products"})
	}
	return c.JSON(products)
}

// GetProductById retrieves a product by its ID.
func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	// Validate ID format
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found!"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve product"})
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var product Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if result := DB.Create(&product); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create product"})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}
