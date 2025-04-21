package routes

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/handlers"
	"Printer3DCourses/internal/middlewares"
	"Printer3DCourses/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, userService *services.UserService, courseService *services.CourseService) {
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
	app.Get("/api/auth/logout", handlers.Logout())
	app.Patch("/api/auth/change-email", middlewares.RequireAuth(cfg), handlers.ChangeEmail(userService))
	app.Patch("/api/auth/change-password", middlewares.RequireAuth(cfg), handlers.ChangePassword(userService))
	app.Patch("/api/auth/change-name", middlewares.RequireAuth(cfg), handlers.ChangeName(userService))
	app.Delete("/api/auth/delete-account", middlewares.RequireAuth(cfg), handlers.DeleteAccount(userService))

	//Роуты для фронта
	app.Get("/", handlers.IndexPage(cfg, courseService))
	app.Get("/starter-kit", handlers.StarterKitPage(cfg))
	app.Get("/profile", middlewares.RequireAuth(cfg), handlers.ProfilePage(cfg, courseService))
	app.Get("/test", handlers.TestingPage(cfg))
	app.Get("/course/:id", handlers.CourseViewPage(cfg))
	app.Get("/homework", handlers.HomeworkPage(cfg))
}
