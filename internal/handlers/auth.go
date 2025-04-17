package handlers

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/dto"
	"Printer3DCourses/internal/models"
	"Printer3DCourses/internal/repositories"
	"Printer3DCourses/internal/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func Registration(db *gorm.DB, cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		input := dto.RegistrationResponse{}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		_, err := repositories.GetUserByEmail(db, input.Email)
		if err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not register user"})
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

		newUser := &models.User{
			LastName:    input.LastName,
			FirstName:   input.FirstName,
			Patronymic:  input.Patronymic,
			Email:       input.Email,
			PhoneNumber: input.PhoneNumber,
			Password:    string(hashedPassword),
		}

		err = repositories.CreateUser(db, newUser)
		if err != nil {
			return err
		}

		token, _ := utils.GenerateJWT(newUser, cfg)

		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "Strict",
		})
		return c.JSON("гуд")
	}
}

func Login(ctx *fiber.Ctx) error {
	return nil
}
