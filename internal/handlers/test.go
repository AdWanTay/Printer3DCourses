package handlers

import (
	"Printer3DCourses/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetTests(testService *services.TestService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		response, err := testService.GetTestByIdForResponse(uint(id))
		if err != nil {
			return err
		}
		return c.JSON(response)
	}
}
