package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebastian009w/go_rest/app/controllers/posts"
)

func PostRoutes(app *fiber.App) {
	app.Get("/posts", posts.Index)
	app.Post("/posts", posts.Create)
}
