package posts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebastian009w/go_rest/database"
	"github.com/sebastian009w/go_rest/models"
)

func Index(ctx *fiber.Ctx) error {

	posts := []models.Post{}

	database.Conn.Find(&posts)

	return ctx.JSON(posts)
}
