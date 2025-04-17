package routes

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/handlers"
	"Printer3DCourses/internal/services"
	"github.com/gofiber/fiber/v2"
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
	//app.Post("/api/auth/login", handlers.Login)
	app.Post("/api/auth/register", handlers.Registration(userService, cfg))

	//Роуты для фронта
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
}
