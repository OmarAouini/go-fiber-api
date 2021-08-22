package main

import (
	"fmt"
	"log"

	"github.com/OmarAouini/go-fiber-api/config"
	"github.com/OmarAouini/go-fiber-api/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//app
	app := fiber.New()

	//middlewares
	app.Use(logger.New(logger.ConfigDefault))
	app.Use(recover.New())
	app.Use(helmet.New())

	//routes mapping
	//config.ConfigRoutes(app)

	//database connect
	config.ConnectDatabase()

	dsn := "user:password@tcp(127.0.0.1)/library?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var userz user.User
	db.First(&userz) // find product with integer primary key

	fmt.Println(userz)

	//404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// content-type check only json allowed
	app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.Status(400).SendString("Only JSON allowed!")
	})

	println("\n")
	fmt.Println(` ██████╗ ██████╗ ██████╗ ███████╗         █████╗ ██████╗ ██╗
██╔════╝██╔═══██╗██╔══██╗██╔════╝        ██╔══██╗██╔══██╗██║
██║     ██║   ██║██████╔╝█████╗          ███████║██████╔╝██║
██║     ██║   ██║██╔══██╗██╔══╝          ██╔══██║██╔═══╝ ██║
╚██████╗╚██████╔╝██║  ██║███████╗███████╗██║  ██║██║     ██║
 ╚═════╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝
	 		ENV:
			VERSION:
	 `)

	//server start
	log.Fatal(app.Listen("0.0.0.0:8080"))

}
