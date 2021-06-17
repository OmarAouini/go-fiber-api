package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	//app
	app := fiber.New()

	//middlewares
	app.Use(logger.New(logger.ConfigDefault))

	//routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello go fiber!")
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("hello, %s", c.Params("name"))
		return c.SendString(msg)
	})

	//404 handler
	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404)
	})

	//server start
	app.Listen(":3000")
}
