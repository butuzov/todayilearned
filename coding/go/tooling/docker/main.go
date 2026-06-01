package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

const AppPort = "8080"

func setupApp() *fiber.App {
	app := fiber.New()

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("ok")
	})

	return app
}

func port() string {
	port := os.Getenv("APP_PORT")
	if port != "" {
		return port
	}
	return AppPort
}

func main() {
	app := setupApp()
	if err := app.Listen(":" + port()); err != nil {
		log.Fatal(err)
	}
}
