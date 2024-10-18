package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qosdil/x_clone_post_svc_gofiber/configs"
)

func main() {
	// Load environment variables
	configs.LoadEnv()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + configs.GetEnv("HTTP_PORT"))
}
