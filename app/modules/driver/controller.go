package driver

import "github.com/gofiber/fiber/v2"

func Controller() *fiber.App {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		
		


		return c.SendString("Hello, World 👋!")
	})

	return app
}
