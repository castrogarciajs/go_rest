package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"Root":  "Ruta inicial",
			"Users": "http://localhost:3000",
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "usuarios desde el backend",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server succesfuly in port 3000")
}
