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
)

func App() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}

	db, err := database.GetConnection(cfg.Database)
	if err != nil {
		return fmt.Errorf("get database connection: %w", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	engine := html.New("./web/templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	routes.SetupRoutes(app, cfg, userService)
	err = app.Listen(":" + cfg.Port)
	if err != nil {
		return fmt.Errorf("app listen: %w", err)
	}
	return nil
}
