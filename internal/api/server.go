// Package api
package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest/handlers"
	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config configs.AppConfig) {
	log.Println("config DSN: ", config.Dsn)
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection error ")
		os.Exit(0)
	}
	db.AutoMigrate(&domain.User{})
	log.Println("db connection: ", db)

	app := fiber.New()
	app.Get("/health", HealthCheck)

	//	handlers.SetupUserRoutes(&restHandler)
	setupRoutes(&rest.RestHandler{
		app,
		db,
	})
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
