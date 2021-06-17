package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func main() {

	//app
	app := fiber.New()

	//middlewares
	app.Use(logger.New(logger.ConfigDefault))
	app.Use(recover.New())
	app.Use(helmet.New())

	//routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello go fiber!")
	})

	app.Post("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("hello, %s", c.Params("name"))
		return c.SendString(msg)
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": "true", "message": "messaggio"})
	})

	//404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	//server start
	log.Fatal(app.Listen(":3000"))
}
