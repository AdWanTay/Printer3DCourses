package handlers

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/notifications"
	"github.com/gofiber/fiber/v2"
)

func StarterKitRequest(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body dto.StarterKitRequest
		if err := c.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Некорректный формат запроса")
		}
		notifications.HandleUserSubmission(cfg, body.FullName, body.Email, body.PhoneNumber)
		return c.JSON(fiber.Map{"message": "Заявка оставлена"})
	}
}
