package middlewares

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/services"
	"Printer3DCourses/internal/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

func RequireAuth(cfg *config.Config, onlyUserId bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token")
		if tokenString == "" && !onlyUserId {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userId, err := utils.ParseAndValidateJWT(tokenString, cfg)
		if err != nil && !onlyUserId {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		if err == nil {
			c.Locals("userId", userId)
		}
		return c.Next()
	}
}

func RequireBought(cfg *config.Config, service *services.UsersCourseService, param string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token")
		if tokenString == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		userId, err := utils.ParseAndValidateJWT(tokenString, cfg)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Invalid courseId")
		}

		isBought, err := service.IsBought(c.Context(), userId, uint(id), param)
		if !isBought {
			return fiber.NewError(fiber.StatusNotAcceptable, "Нет доступа к ресурсу")
		}

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Next()
	}
}
