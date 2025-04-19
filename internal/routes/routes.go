package routes

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/handlers"
	"Printer3DCourses/internal/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, userService *services.UserService) {
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		return c.Next()
	})

	//app.Static(config.UploadsURL, "."+config.UploadsURL)
	app.Static("/web", "./web", fiber.Static{CacheDuration: 0})

	//Роуты для апи
	app.Post("/api/auth/login", handlers.Login(userService, cfg))
	app.Post("/api/auth/register", handlers.Registration(userService, cfg))
	app.Get("/api/auth/logout", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour), // Ставим просроченное время
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		return c.JSON(fiber.Map{"message": "Вы успешно вышли из аккаунта"})
	})

	//Роуты для фронта
	app.Get("/", handlers.IndexPage(cfg))
	app.Get("/starter-kit", handlers.StarterKitPage(cfg))
	app.Get("/profile", handlers.ProfilePage(cfg))
	app.Get("/test", handlers.TestingPage(cfg))
	app.Get("/course/:id", handlers.CourseViewPage(cfg))
}
