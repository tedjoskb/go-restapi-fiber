package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedjoskb/go-restapi-fiber/database"
	"github.com/tedjoskb/go-restapi-fiber/database/migration"
	"github.com/tedjoskb/go-restapi-fiber/route"
)

func main() {

	database.ConnectionDatabase()
	migration.RunMigration()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, This is Fiber!")
	})

	route.RouteInit(app)

	app.Listen(":3000")

}
