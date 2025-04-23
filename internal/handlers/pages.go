package handlers

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/services"
	"Printer3DCourses/internal/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Render(c *fiber.Ctx, template string, data fiber.Map, cfg *config.Config, userService *services.UserService) error {
	if data == nil {
		data = fiber.Map{}
	}

	tokenString := c.Cookies("token")
	userId, err := utils.ParseAndValidateJWT(tokenString, cfg)

	if err == nil {
		data["firstName"], data["lastName"], _ = userService.GetFirstNameAndLastName(c.Context(), userId)
	}

	data["isAuthenticated"] = err == nil

	return c.Render(template, data)
}

func IndexPage(cfg *config.Config, courseService *services.CourseService, userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userIdPointer *uint
		if userId, ok := c.Locals("userId").(uint); ok {
			userIdPointer = &userId
		}
		response, err := courseService.GetCoursesForResponse(c.Context(), userIdPointer)
		if err != nil {
			return err
		}

		return Render(c, "index", fiber.Map{
			"items": response.Items,
		}, cfg, userService)
	}
}
func StarterKitPage(cfg *config.Config, userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "starter-kit", nil, cfg, userService)
	}
}
func ProfilePage(cfg *config.Config, courseService *services.CourseService, userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(uint)
		usersCoursesResponse, err := courseService.GetAllPaidCoursesForResponse(c.Context(), userId)
		if err != nil {
			return err
		}
		return Render(c, "profile", fiber.Map{"response": usersCoursesResponse, "page": "Profile"}, cfg, userService)
	}
}
func TestingPage(cfg *config.Config, userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "test/test", nil, cfg, userService)
	}
}
func CourseViewPage(cfg *config.Config, userService *services.UserService, courseService *services.CourseService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(uint)
		courseId, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		coursePageResponse, err := courseService.GetCourseDetailInfo(c.Context(), userId, uint(courseId))

		return Render(c, "course-view", fiber.Map{"response": coursePageResponse}, cfg, userService)
	}
}
func HomeworkPage(cfg *config.Config, userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return Render(c, "homework", nil, cfg, userService)
	}
}
