package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Listen(":3000")
	fmt.Println("Server succesfuly in port 3000")
}
