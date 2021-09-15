package main

import (
	"github.com/OmarAouini/go-fiber-api/user"
	"github.com/gofiber/fiber/v2"
)

//function used to config api routes,
//add here new api endpoints
func ConfigRoutes(app *fiber.App) {

	// content-type check only json allowed
	// app.Use(func(c *fiber.Ctx) error {
	// 	if c.Is("json") {
	// 		return c.Next()
	// 	}
	// 	return c.Status(400).SendString("Only JSON allowed!")
	// })

	//ROUTES

	//health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	//root
	api := app.Group("/api/v1")

	//users
	users := api.Group("/users")
	users.Get("/", user.Get_all)
	users.Get("/name/:username", user.Find)
	users.Post("/create", user.Create)
	users.Put("/update/:name", user.Update)
	users.Delete("/delete/:name", user.Delete)

	//products
	//products := api.Group("/products")
	//products.Get("/", prodotti.Get_all)
}
