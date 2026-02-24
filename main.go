package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
)

func main() {
	fmt.Println("teste")
	fmt.Println("2222")

	configs.LoadAppSettings()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("for testing only")
	})
	app.Listen(":9000")
}
