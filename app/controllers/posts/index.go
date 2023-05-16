package posts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebastian009w/go_rest/app/models"
	"github.com/sebastian009w/go_rest/platform/database"
)

func Index(ctx *fiber.Ctx) error {

	posts := []models.Post{}

	database.Conn.Find(&posts)

	return ctx.JSON(posts)
}
