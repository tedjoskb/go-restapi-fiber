package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedjoskb/go-restapi-fiber/config"
	"github.com/tedjoskb/go-restapi-fiber/controllers"
	"github.com/tedjoskb/go-restapi-fiber/middleware"
)

func RouteInit(r *fiber.App) {

	r.Static("/public", config.ProjectRootPath+"/public/asset") //route folder asset (utk gambar,file dll)

	// /api/
	api := r.Group("/api")
	api.Post("/login", middleware.Auth, controllers.Login)
	// /api/user/
	users := api.Group("/user")
	users.Get("/", middleware.Auth, controllers.GetUserAll)
	users.Get("/:id", middleware.Auth, controllers.GetUserById)
	users.Post("/", middleware.Auth, controllers.CreateUser)
	users.Put("/:id", middleware.Auth, controllers.UpdateUser)
	users.Post("/update-multiple", middleware.Auth, controllers.UpdateMultipleUsers)
	users.Post("/delete-user", middleware.Auth, controllers.SoftDeleteUser)

	book := api.Group("/books")
	book.Get("/", middleware.Auth, controllers.Index)
	book.Get("/:id", middleware.Auth, controllers.Show)
	book.Post("/", middleware.Auth, controllers.Create)
	book.Put("/:id", middleware.Auth, controllers.Update)
	book.Delete("/:id", middleware.Auth, controllers.Delete)

}
