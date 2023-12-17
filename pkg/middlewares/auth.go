package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/config"
	"github.com/santoadji21/santoadji21-go-fiber-product-api/pkg/utils"
)

func Protected() fiber.Handler {
    jwtCfg := config.JwtCfg()	
    return jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(jwtCfg.SecretKey)},
        ErrorHandler: jwtError,
    })
}

func jwtError(c *fiber.Ctx, err error) error {
    return c.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{
		Success: false,
		Message: "Unauthorized",
		Data:    nil,
	})
}
