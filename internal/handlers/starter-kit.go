package handlers

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/notifications"
	"github.com/gofiber/fiber/v2"
)

func StarterKitRequest(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		notifications.HandleUserSubmission(cfg)
		return c.JSON(fiber.Map{"message": "Заявка оставлена"})
	}
}
