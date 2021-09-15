package main

import (
	"log"

	"github.com/OmarAouini/go-fiber-api/database"
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
	ConfigRoutes(app)

	//database connect
	database.InitDatabase()

	//404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	//server start
	log.Fatal(app.Listen("0.0.0.0:8080"))
}
