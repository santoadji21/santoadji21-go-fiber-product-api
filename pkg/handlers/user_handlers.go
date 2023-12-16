package handlers

import (
	"go-product-api/internal/db"
	"go-product-api/pkg/models"
	"go-product-api/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
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

    result := db.GetDB().Create(user)
	
     // Check for duplicate email error
   if result.Error != nil {
        // Check for duplicate email error using string matching
        if strings.Contains(result.Error.Error(), "users_email_key") {
            return c.Status(fiber.StatusConflict).JSON(utils.ApiResponse{
                Success: false,
                Message: "Email already in use",
                Data:    nil,
            })
        }

        // Handle other potential errors
        return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
            Success: false,
            Message: "Failed to create user",
            Data:    result.Error.Error(),
        })
    }


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
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status":  "error",
            "message": "User not found",
            "error":   result.Error.Error(),
        })
    }

    return c.JSON(user)
}


func UpdateUser(c *fiber.Ctx) error {
    userID := c.Params("id")
    var user models.User

    if err := db.GetDB().First(&user, userID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status":  "error",
            "message": "User not found",
            "error":   err.Error(),
        })
    }

    type UpdateUserInput struct {
        FirstName string
        LastName  string
        Email     string
    }
    var input UpdateUserInput

    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": "Failed to parse body",
            "error":   err.Error(),
        })
    }

    db.GetDB().Model(&user).Updates(models.User{FirstName: input.FirstName, LastName: input.LastName, Email: input.Email})

    return c.JSON(user)
}


func DeleteUser(c *fiber.Ctx) error {
    userID := c.Params("id")

    result := db.GetDB().Delete(&models.User{}, userID)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": "Failed to delete user",
            "error":   result.Error.Error(),
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}

