package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// CreateProduct - Handler for creating a new product
func CreateProduct(c *fiber.Ctx) error {
    return c.SendString("CreateProduct")
}

// GetProduct - Handler for getting a product's details
func GetProduct(c *fiber.Ctx) error {
   return c.SendString("GetProduct")
}

// UpdateProduct - Handler for updating a product's details
func UpdateProduct(c *fiber.Ctx) error {
    return c.SendString("UpdateProduct")
}

// DeleteProduct - Handler for deleting a product
func DeleteProduct(c *fiber.Ctx) error {
    return c.SendString("DeleteProduct")
}
