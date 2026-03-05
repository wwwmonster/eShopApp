// Package api
package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()
	log.Println("------fiber server started")
	//	var err error

	app.Get("/health", HealthCheck)
	if err := app.Listen(config.ServerPort); err != nil {
		os.Exit(0)
	}
}

func HealthCheck(ctx *fiber.Ctx) error {
	log.Println("health check")
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "This is health check breathing...",
	})
}
