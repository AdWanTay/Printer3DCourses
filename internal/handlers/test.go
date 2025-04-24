package handlers

import (
	"Printer3DCourses/internal/dto"
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

func SaveResult(testService *services.TestService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var body dto.TestAnswersRequest
		if err := c.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Некорректный формат запроса")
		}

		userId := c.Locals("userId").(uint)

		result, err := testService.SaveTestResult(c.Context(), userId, &body)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Ошибка сохранения теста")
		}

		return c.JSON(fiber.Map{"result": int(result * 100)})
	}
}
