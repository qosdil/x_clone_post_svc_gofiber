package main

import (
	"context"
	"x_clone_post_svc_gofiber/configs"
	db "x_clone_post_svc_gofiber/repository/databases"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	configs.LoadEnv()

	// Get repository
	repo := db.NewDummyRepository()

	app := fiber.New()
	ctx := context.Background()
	app.Get("/v1/posts", func(c *fiber.Ctx) error {
		posts, err := repo.Find(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"posts": posts})
	})
	app.Get("/v1/posts/:id?", func(c *fiber.Ctx) error {
		post, err := repo.FirstByID(ctx, c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"post": post})
	})

	app.Listen(":" + configs.GetEnv("HTTP_PORT"))
}
