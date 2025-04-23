package routes

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/handlers"
	"Printer3DCourses/internal/middlewares"
	"Printer3DCourses/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, userService *services.UserService, courseService *services.CourseService, testService *services.TestService) {
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
	app.Patch("/api/auth/change-email", middlewares.RequireAuth(cfg, false), handlers.ChangeEmail(userService))
	app.Patch("/api/auth/change-password", middlewares.RequireAuth(cfg, false), handlers.ChangePassword(userService))
	app.Patch("/api/auth/change-name", middlewares.RequireAuth(cfg, false), handlers.ChangeName(userService))
	app.Patch("/api/auth/change-phone", middlewares.RequireAuth(cfg, false), handlers.ChangePhoneNumber(userService))
	app.Delete("/api/auth/delete-account", middlewares.RequireAuth(cfg, false), handlers.DeleteAccount(userService))

	app.Get("/api/tests/:id", middlewares.RequireAuth(cfg, false), handlers.GetTests(testService))

	//Роуты для фронта
	app.Get("/", middlewares.RequireAuth(cfg, true), handlers.IndexPage(cfg, courseService, userService))
	app.Get("/starter-kit", handlers.StarterKitPage(cfg, userService))
	app.Get("/profile", middlewares.RequireAuth(cfg, false), handlers.ProfilePage(cfg, courseService, userService))
	app.Get("/test", handlers.TestingPage(cfg, userService))
	app.Get("/course/:id", middlewares.RequireAuth(cfg, false), handlers.CourseViewPage(cfg, userService, courseService))
	app.Get("/homework", handlers.HomeworkPage(cfg, userService))
}
