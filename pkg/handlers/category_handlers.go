package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/internal/db"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/models"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/utils"
)

// CreateCategory - Handler for creating a new category
func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success: false,
			Message: "Error parsing JSON",
			Data:    err.Error(),
		})
	}

	// Check if a category with the same name already exists
	var existingCategory models.Category
	result := db.GetDB().Where("name = ?", category.Name).First(&existingCategory)
	if result.Error == nil {
		// A category with the same name was found
		return c.Status(fiber.StatusConflict).JSON(utils.ApiResponse{
			Success: false,
			Message: "Category name already exists",
			Data:    nil,
		})
	}

	// No existing category found, proceed to create a new one
	result = db.GetDB().Create(&category)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to create category",
			Data:    result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Success: true,
		Message: "Category created successfully",
		Data:    category,
	})
}

// GetAllCategories - Handler for getting all categories
func GetAllCategories(c *fiber.Ctx) error {
	var categories []models.Category
	result := db.GetDB().Find(&categories)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to retrieve categories",
			Data:    result.Error.Error(),
		})
	}

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Categories retrieved successfully",
		Data:    categories,
	})
}

// GetCategory - Handler for getting a category's details
func GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	result := db.GetDB().First(&category, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "Category not found",
			Data:    nil,
		})
	}

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Category retrieved successfully",
		Data:    category,
	})
}

// UpdateCategory - Handler for updating a category's details
func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := db.GetDB().First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "Category not found",
			Data:    nil,
		})
	}

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success: false,
			Message: "Error parsing JSON",
			Data:    err.Error(),
		})
	}

	db.GetDB().Save(&category)

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Category updated successfully",
		Data:    category,
	})
}

// DeleteCategory - Handler for deleting a category
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	result := db.GetDB().Delete(&models.Category{}, id)
	if err := db.GetDB().First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "Category not found",
			Data:    nil,
		})
	}

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to delete category",
			Data:    result.Error.Error(),
		})
	}

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Category deleted successfully",
		Data:    nil,
	})
}
