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

func Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour), // Ставим просроченное время
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})
		return c.JSON(fiber.Map{"message": "Вы успешно вышли из аккаунта"})
	}
}

func ChangeEmail(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			NewEmail string `json:"new_email"`
		}

		var body request
		if err := c.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Некорректный формат запроса")
		}

		if body.NewEmail == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Новый email не может быть пустым")
		}

		userID := c.Locals("userId").(uint)

		if err := userService.ChangeEmail(c.Context(), userID, body.NewEmail); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Не удалось изменить email")
		}

		return c.JSON(fiber.Map{"message": "Email успешно изменён"})
	}
}

func ChangePassword(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			NewPassword string `json:"new_password"`
		}

		var body request
		if err := c.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Некорректный формат запроса")
		}

		if body.NewPassword == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Новый пароль не может быть пустым")
		}

		userID := c.Locals("userId").(uint)

		if err := userService.ChangePassword(c.Context(), userID, body.NewPassword); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Не удалось изменить пароль")
		}

		return c.JSON(fiber.Map{"message": "Пароль успешно изменён"})
	}
}

func ChangeName(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			NewLastName   string `json:"new_last_name"`
			NewFirstName  string `json:"new_first_name"`
			NewPatronymic string `json:"new_patronymic"`
		}

		var body request
		if err := c.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Некорректный формат запроса")
		}

		if body.NewFirstName == "" || body.NewLastName == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Фамилия и имя не могут быть пустыми")
		}

		userID := c.Locals("userId").(uint)

		if err := userService.ChangeName(c.Context(), userID, body.NewLastName, body.NewFirstName, body.NewPatronymic); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Не удалось изменить ФИО")
		}

		return c.JSON(fiber.Map{"message": "ФИО успешно изменены"})
	}
}

func DeleteAccount(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userId").(uint)

		if err := userService.DeleteUser(c.Context(), userID); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Внутренняя ошибка сервера")
		}

		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})
		return c.JSON(fiber.Map{"message": "Аккаунт успешно удален"})
	}
}

func ChangePhoneNumber(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			NewPhoneNumber string `json:"new_phone_number"`
		}

		var body request
		if err := c.BodyParser(&body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Некорректный формат запроса")
		}

		if body.NewPhoneNumber == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Номер телефона не может быть пустым")
		}

		userID := c.Locals("userId").(uint)

		if err := userService.ChangePhoneNumber(c.Context(), userID, body.NewPhoneNumber); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Не удалось обновить номер телефона")
		}

		return c.JSON(fiber.Map{"message": "номер телефона успешно изменен"})
	}
}
