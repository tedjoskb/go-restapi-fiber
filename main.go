package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedjoskb/go-restapi-fiber/controllers/bookcontroller"
	"github.com/tedjoskb/go-restapi-fiber/database"
	"github.com/tedjoskb/go-restapi-fiber/database/migration"
)

func main() {
	database.ConnectionDatabase()
	migration.RunMigration()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, This is Fiber!")
	})

	// /api/books/1
	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("/", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

	app.Listen(":3000")

}
