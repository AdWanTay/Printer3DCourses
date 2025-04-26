package main

import (
	"Printer3DCourses/internal"
	"Printer3DCourses/internal/config"
	"github.com/gofiber/fiber/v2/log"
)

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
