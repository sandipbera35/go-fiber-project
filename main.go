package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Data struct {
	Name  string
	Email string
}

func main() {
	fmt.Println("Demo Project")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		var v1 Data
		v1.Name = "Sandip"
		v1.Email = "s@b.com"
		return c.JSON(v1)
		// return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}
