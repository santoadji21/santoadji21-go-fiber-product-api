package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/internal/db"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/models"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/utils"
)

// CreateProduct - Handler for creating a new product
// @Summary Create a new product
// @Description Create a new product with the given details
// @Tags Product
// @Accept  json
// @Produce  json
// @Param   products body     models.Product   true  "Product Info"
// @Success 201 {object}  models.Product
// @Router /api/product [post]
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success: false,
			Message: "Error parsing JSON",
			Data:    err.Error(),
		})
	}

	// Check if a product with the same name already exists
	var existingProduct models.Product
	result := db.GetDB().Where("name = ?", product.Name).First(&existingProduct)
	if result.Error == nil {
		// A product with the same name was found
		return c.Status(fiber.StatusConflict).JSON(utils.ApiResponse{
			Success: false,
			Message: "Product name already exists",
			Data:    nil,
		})
	}

	// No existing product found, proceed to create a new one
	result = db.GetDB().Create(&product)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to create product",
			Data:    result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Success: true,
		Message: "Product created successfully",
		Data:    product,
	})
}

// GetAllProducts - Handler for getting all products
// GetAllProducts retrieves all products
// @Summary Get all products
// @Description Retrieves a list of all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /api/products [get]
func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	result := db.GetDB().Find(&products)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to retrieve products",
			Data:    result.Error.Error(),
		})
	}

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

// GetProduct - Handler for getting a product's details
// GetProduct retrieves a single product by ID
// @Summary Get a product
// @Description Retrieves a product by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} utils.ApiResponse "Product not found"
// @Router /api/product/{id} [get]
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	result := db.GetDB().First(&product, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "Product not found",
			Data:    nil,
		})
	}

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Product retrieved successfully",
		Data:    product,
	})
}

// UpdateProduct - Handler for updating a product's details
// UpdateProduct updates a product's details
// @Summary Update a product
// @Description Updates a product's details by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product update data"
// @Success 200 {object} models.Product
// @Failure 404 {object} utils.ApiResponse "Product not found"
// @Router /api/product/{id} [patch] update product
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := db.GetDB().First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "Product not found",
			Data:    nil,
		})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success: false,
			Message: "Error parsing JSON",
			Data:    err.Error(),
		})
	}

	db.GetDB().Save(&product)

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Product updated successfully",
		Data:    product,
	})
}

// DeleteProduct - Handler for deleting a product
// DeleteProduct deletes a product
// @Summary Delete a product
// @Description Deletes a product by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse "Product not found"
// @Router /api/product/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	result := db.GetDB().Delete(&models.Product{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to delete product",
			Data:    result.Error.Error(),
		})
	}

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Product deleted successfully",
		Data:    nil,
	})
}
