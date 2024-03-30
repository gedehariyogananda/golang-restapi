package routes

import (
	"test/golang/controller"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {

	// grouping api routes
	api := app.Group("/api")

	// first of all books api
	books := api.Group("/books")

	// routes books api
	books.Get("/", bookcontroller.Index)
	books.Get("/:id", bookcontroller.Show)
	books.Post("/", bookcontroller.Store)
	books.Put("/:id", bookcontroller.Update)
	books.Delete("/:id", bookcontroller.Delete)
}
