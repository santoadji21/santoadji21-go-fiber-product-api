package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// CreateCategory - Handler for creating a new category
func CreateCategory(c *fiber.Ctx) error {
    return c.SendString("CreateCategory")
}

// GetCategory - Handler for getting a category's details
func GetCategory(c *fiber.Ctx) error {
   return c.SendString("GetCategory")
}

// UpdateCategory - Handler for updating a category's details
func UpdateCategory(c *fiber.Ctx) error {
    return c.SendString("UpdateCategory")
}

// DeleteCategory - Handler for deleting a category
func DeleteCategory(c *fiber.Ctx) error {
    return c.SendString("DeleteCategory")
}
