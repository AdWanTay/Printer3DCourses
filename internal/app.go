package internal

import (
	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		return c.Next()
	})

	//app.Static(config.UploadsURL, "."+config.UploadsURL)
	app.Static("/web", "./web", fiber.Static{CacheDuration: 0})

	//Роуты для апи

	//Роуты для фронта
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
}

func App() {

	engine := html.New("./web/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoutes(app)
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
