package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/config"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/internal/db"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/models"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/utils"
)

// Login authenticates a user and provides a JWT token
// @Summary User login
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   email    body    string  true  "Email address of the user"
// @Param   password body    string  true  "Password of the user"
// @Success 200 {object} map[string]interface{} "Login successful with token"
// @Failure 400 {object} utils.ApiResponse "Invalid input"
// @Failure 401 {object} utils.ApiResponse "Unauthorized"
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /api/login [post]
func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Success: false,
			Message: "Cannot parse JSON",
			Data:    err.Error(),
		})
	}

	var user models.User
	result := db.GetDB().Where("email = ?", request.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ApiResponse{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
	}

	if !utils.ValidatePassword(request.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{
			Success: false,
			Message: "Incorrect password",
			Data:    nil,
		})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	jwtCfg := config.JwtCfg()
	t, err := token.SignedString([]byte(jwtCfg.SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating token")
	}

	return c.JSON(utils.ApiResponse{
		Success: true,
		Message: "Login successful",
		Data:    fiber.Map{"token": t},
	})
}
