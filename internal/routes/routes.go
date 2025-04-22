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
	app.Patch("/api/auth/change-email", middlewares.RequireAuth(cfg), handlers.ChangeEmail(userService))
	app.Patch("/api/auth/change-password", middlewares.RequireAuth(cfg), handlers.ChangePassword(userService))
	app.Patch("/api/auth/change-name", middlewares.RequireAuth(cfg), handlers.ChangeName(userService))
	//TODO change-phone
	app.Patch("/api/auth/change-phone", middlewares.RequireAuth(cfg), handlers.ChangePhoneNumber(userService))
	app.Delete("/api/auth/delete-account", middlewares.RequireAuth(cfg), handlers.DeleteAccount(userService))

	app.Get("/api/tests/:id", middlewares.RequireAuth(cfg), handlers.GetTests(testService))

	//Роуты для фронта
	app.Get("/", handlers.IndexPage(cfg, courseService, userService))
	app.Get("/starter-kit", handlers.StarterKitPage(cfg, userService))
	app.Get("/profile", middlewares.RequireAuth(cfg), handlers.ProfilePage(cfg, courseService, userService))
	app.Get("/test", handlers.TestingPage(cfg, userService))
	app.Get("/course/:id", handlers.CourseViewPage(cfg, userService))
	app.Get("/homework", handlers.HomeworkPage(cfg, userService))
}
