package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New(fiber.Config{
		AppName: "Test App v1.0.1",
	})

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Hello from golang",
		})
	})

	app.Listen(":3000")

}
