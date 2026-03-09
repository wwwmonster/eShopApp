// Package api
package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest/handlers"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()
	log.Println("------fiber server started")
	//	var err error

	app.Get("/health", HealthCheck)

	restHandler := &rest.RestHandler{
		app,
	}
	fmt.Println("==============")
	//	handlers.SetupUserRoutes(&restHandler)
	setupRoutes(restHandler)
	if err := app.Listen(config.ServerPort); err != nil {
		os.Exit(0)
	}
}

func setupRoutes(rh *rest.RestHandler) {
	// user handlers
	handlers.SetupUserRoutes(rh)

	// transactions

	// catalog
}

func HealthCheck(ctx *fiber.Ctx) error {
	log.Println("health check")
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "This is health check, breathing...",
	})
}
