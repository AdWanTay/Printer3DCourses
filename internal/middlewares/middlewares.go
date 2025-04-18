package middlewares

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func RequireAuth(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": "Missing token"})
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userId, err := utils.ParseAndValidateJWT(tokenString, cfg)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": "Invalid token"})
		}

		c.Locals("user_id", userId)
		return c.Next()
	}
}
