package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/internal/db"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/models"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to parse JSON body",
			Data:    err.Error(),
		})
	}

	// Check if a user with the same email already exists
	var existingUser models.User
	result := db.GetDB().Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		// A user with the same email was found
		return c.Status(fiber.StatusConflict).JSON(utils.ApiResponse{
			Success: false,
			Message: "Email already in use",
			Data:    nil,
		})
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to hash password",
			Data:    err.Error(),
		})
	}
	user.Password = string(hash)

	// Create the user
	result = db.GetDB().Create(user)
	if result.Error != nil {
		// Handle other potential errors
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to create user",
			Data:    result.Error.Error(),
		})
	}

	// Exclude the password from the response
	user.Password = ""
	return c.Status(fiber.StatusCreated).JSON(utils.ApiResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	result := db.GetDB().Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to query users",
			Data:    result.Error.Error(),
		})
	}

	// Set the password field to an empty string for each user
	for i := range users {
		users[i].Password = ""
	}

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user models.User

	result := db.GetDB().First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "User not found",
			Data:    result.Error.Error(),
		})
	}

	user.Password = ""
	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user models.User

	if err := db.GetDB().First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
	}

	type UpdateUserInput struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email" validate:"email"`
	}
	var input UpdateUserInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success: false,
			Message: "Invalid payload",
			Data:    err.Error(),
		})
	}

	db.GetDB().Model(&user).Updates(models.User{FirstName: input.FirstName, LastName: input.LastName, Email: input.Email})
	user.Password = ""
	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user models.User
	result := db.GetDB().Delete(&models.User{}, userID)

	if err := db.GetDB().First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
	}

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Success: false,
			Message: "Failed to delete user",
			Data:    result.Error.Error(),
		})
	}

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "User deleted successfully",
		Data:    nil,
	})
}
