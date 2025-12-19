package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pranaydhanke/go-user-api/internal/handler"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	app.Get("/users", h.List)          // GET /users
	app.Post("/users", h.Create)       // POST /users
	app.Get("/users/:id", h.Get)       // GET /users/1
	app.Put("/users/:id", h.Update)    // PUT /users/1
	app.Delete("/users/:id", h.Delete) // DELETE /users/1
}
