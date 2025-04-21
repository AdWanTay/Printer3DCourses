package handlers

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/services"
	"Printer3DCourses/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, template string, data fiber.Map, cfg *config.Config) error {
	if data == nil {
		data = fiber.Map{}
	}

	tokenString := c.Cookies("token")
	_, err := utils.ParseAndValidateJWT(tokenString, cfg)

	data["isAuthenticated"] = err == nil
	data["firstName"] = "Вася"
	data["lastName"] = "Пупкин"
	return c.Render(template, data)
}

func IndexPage(cfg *config.Config, courseService *services.CourseService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		response, err := courseService.GetCoursesForResponse(c.Context())
		if err != nil {
			return err
		}

		return Render(c, "index", fiber.Map{
			"items": response.Items,
		}, cfg)
	}
}
func StarterKitPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "starter-kit", nil, cfg)
	}
}
func ProfilePage(cfg *config.Config, courseService *services.CourseService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(uint)
		usersCoursesResponse, err := courseService.GetAllPaidCoursesForResponse(c.Context(), userId)
		if err != nil {
			return err
		}
		return Render(c, "profile", fiber.Map{"response": usersCoursesResponse, "page": "Profile"}, cfg)
	}
}
func TestingPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "test/test", nil, cfg)
	}
}
func CourseViewPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "course-view", nil, cfg)
	}
}
func HomeworkPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "homework", nil, cfg)
	}
}
