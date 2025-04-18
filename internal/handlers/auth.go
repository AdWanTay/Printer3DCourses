package handlers

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/services"
	"Printer3DCourses/internal/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

func Registration(s *services.UserService, cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.RegistrationRequest

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		user, err := s.RegisterUser(c.Context(), input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		token, _ := utils.GenerateJWT(user, cfg)

		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "Strict",
		})

		return c.JSON(fiber.Map{"message": "Registration successful"})
	}
}

func Login(s *services.UserService, cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.LoginRequest

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		user, err := s.LoginUser(c.Context(), input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		token, _ := utils.GenerateJWT(user, cfg)

		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "Strict",
		})

		return c.JSON(fiber.Map{"message": "Login successful"})
	}
}
