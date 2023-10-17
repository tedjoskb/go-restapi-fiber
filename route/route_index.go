package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedjoskb/go-restapi-fiber/config"
	"github.com/tedjoskb/go-restapi-fiber/controllers"
)

func RouteInit(r *fiber.App) {

	r.Static("/public", config.ProjectRootPath+"/public/asset") //route folder asset (utk gambar,file dll)

	// /api/
	api := r.Group("/api")
	// /api/user/
	users := api.Group("/user")
	users.Get("/", controllers.GetUserAll)
	users.Get("/:id", controllers.GetUserById)
	users.Post("/", controllers.CreateUser)
	users.Put("/:id", controllers.UpdateUser)
	users.Post("/update-multiple", controllers.UpdateMultipleUsers)
	users.Post("/delete-user", controllers.SoftDeleteUser)

	book := api.Group("/books")
	book.Get("/", controllers.Index)
	book.Get("/:id", controllers.Show)
	book.Post("/", controllers.Create)
	book.Put("/:id", controllers.Update)
	book.Delete("/:id", controllers.Delete)

}
