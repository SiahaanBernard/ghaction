package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/question", func(c *fiber.Ctx) {
		c.Send("Thank you for your question")
	})

	app.Listen(3000)
}
