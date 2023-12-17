package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/internal/db"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/models"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/utils"
)

// CreateCategory - Handler for creating a new category
// CreateCategory creates a new category
// @Summary Create a new category
// @Description Create a new category with the given name
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.Category true "Category Info"
// @Success 201 {object} models.Category
// @Router /api/category [post]
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
// GetAllCategories retrieves all categories
// @Summary Get all categories
// @Description Retrieves a list of all categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {array} models.Category
// @Router /api/categories [get]
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
// GetCategory retrieves a single category by ID
// @Summary Get a category
// @Description Retrieves a category by its ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} utils.ApiResponse "Category not found"
// @Router /api/category/{id} [get]
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
// UpdateCategory updates a category's details
// @Summary Update a category
// @Description Updates a category's details by its ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category update data"
// @Success 200 {object} models.Category
// @Failure 404 {object} utils.ApiResponse "Category not found"
// @Router /api/category/{id} [patch]
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
// DeleteCategory deletes a category
// @Summary Delete a category
// @Description Deletes a category by its ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse "Category not found"
// @Router /api/category/{id} [delete]
func DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success: false,
			Message: "Invalid ID format",
			Data:    nil,
		})
	}

	var category models.Category
	if err := db.GetDB().First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "Category not found",
			Data:    nil,
		})
	}

	result := db.GetDB().Unscoped().Delete(&models.Category{}, id)
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
