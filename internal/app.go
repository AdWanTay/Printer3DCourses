package internal

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/database"
	"Printer3DCourses/internal/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func App() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}

	db, err := database.GetConnection()
	if err != nil {
		return fmt.Errorf("get database connection: %w", err)
	}

	engine := html.New("./web/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.SetupRoutes(app, cfg, db)
	err = app.Listen(":8080")
	if err != nil {
		return fmt.Errorf("app listen: %w", err)
	}
	return nil
}
