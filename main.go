package main

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello %s \n", c.Params("name")))
	})

	app.Listen(":3000")
}
