package main

import (
	"Printer3DCourses/internal"
	"Printer3DCourses/internal/config"
	"github.com/gofiber/fiber/v2/log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("get config: %v", err)
	}

	go func() {
		if err := internal.App(cfg); err != nil {
			log.Fatalf("App error: %v", err)
		}
	}()

	if err := internal.TgApp(cfg); err != nil {
		log.Fatalf("TgApp error: %v", err)
	}

}
