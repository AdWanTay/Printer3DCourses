package internal

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/database"
	"Printer3DCourses/internal/repositories"
	"Printer3DCourses/internal/routes"
	"Printer3DCourses/internal/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"strings"
)

func App(cfg *config.Config) error {

	db, err := database.GetConnection(cfg.Database)
	if err != nil {
		return fmt.Errorf("get database connection: %w", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	testRepo := repositories.NewTestRepository(db)
	testService := services.NewTestService(testRepo)

	courseRepo := repositories.NewCourseRepository(db)
	courseService := services.NewCourseService(courseRepo, userRepo, testRepo)

	engine := html.New("./web/templates", ".html")
	engine.Reload(true)

	engine.AddFunc("split", func(s, sep string) []string {
		return strings.Split(s, sep)
	})

	app := fiber.New(fiber.Config{Views: engine})

	routes.SetupRoutes(app, cfg, userService, courseService, testService)
	err = app.Listen(":" + cfg.Port)
	if err != nil {
		return fmt.Errorf("app listen: %w", err)
	}
	return nil
}
