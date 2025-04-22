package middlewares

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func RequireAuth(cfg *config.Config, onlyUserId bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token")
		if tokenString == "" && !onlyUserId {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": "Missing token"})
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userId, err := utils.ParseAndValidateJWT(tokenString, cfg)
		if err != nil && !onlyUserId {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": "Invalid token"})
		}
		if err == nil {
			c.Locals("userId", userId)
		}
		return c.Next()
	}
}
