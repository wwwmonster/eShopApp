// Package api
package api

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()
	//	var err error

	if err := app.Listen(config.ServerPort); err != nil {
		os.Exit(0)
	}
}
