package main

import (
	"handson-httphandler/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	middleware.AddJWTRoute(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
