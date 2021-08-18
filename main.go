package main

import (
	"log"

	"github.com/OmarAouini/go-fiber-api/config"
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

	//routes mapping
	config.ConfigRoutes(app)

	//404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	//content-type check only json allowed
	app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})

	//server start
	log.Fatal(app.Listen(":8080"))
}
