package handlers

import (
	"Printer3DCourses/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func BuyCourse(usersCourseService *services.UsersCourseService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(uint)
		courseId, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
		}
		err = usersCourseService.BuyCourse(c.Context(), userId, uint(courseId))
		if err != nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "buy course successfully"})
	}
}
