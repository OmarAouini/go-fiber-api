package config

import (
	"github.com/OmarAouini/go-fiber-api/person"
	"github.com/gofiber/fiber/v2"
)

//function used to config routes,
//add here new api endpoints
func ConfigRoutes(app *fiber.App) {

	//ROUTES

	//health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	//root
	api := app.Group("/api/v1")

	//users
	users := api.Group("/users")
	users.Get("/", person.Get_all)
	users.Get("/name/:name", person.Find)
	users.Post("/create", person.Create)
	users.Put("/update/:name", person.Update)
	users.Delete("/delete/:name", person.Delete)
}
