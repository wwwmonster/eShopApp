// Package api
package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest/handlers"
	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"github.com/wwwmonster/eShopApp/go/v2/internal/helper"
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

	if err = db.AutoMigrate(&domain.User{}, &domain.Address{}, &domain.Cart{},
		&domain.Order{}, &domain.BankAccount{},
		&domain.Category{}, &domain.Product{}); err != nil {
		log.Fatal("error on running migration", err.Error())
	}
	log.Println("db connection: ", db)
	auth := helper.SetupAuth(config.AppSecret, 2)

	app := fiber.New()
	app.Get("/health", HealthCheck)

	//	handlers.SetupUserRoutes(&restHandler)
	setupRoutes(&rest.RestHandler{
		App: app, Db: db, ConnPool: GetEmsPgxConnPool(), Auth: auth, Config: config,
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
	handlers.SetupCatalogRoutes(rh)
}

func HealthCheck(ctx *fiber.Ctx) error {
	log.Println("health check")
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "This is health check, breathing...",
	})
}

var emsPgxConnPool *pgxpool.Pool

func CreateConnectionPool() {
	// Parse the connection string into a Config
	var ctx = context.Background()
	// config, err := pgxpool.ParseConfig(os.Getenv("PGCONNSTRING"))
	config, err := pgxpool.ParseConfig("postgres://root:root@localhost:5432/online-shopping")
	if err != nil {
		log.Fatal("unable to parse connection string: %w", err)
		os.Exit(1)
	}

	// Optional: Configure pool settings
	config.MaxConns = 10 // Maximum number of connections
	config.MinConns = 2  // Minimum number of connections
	// config.MaxConnLifetime = time.Hour // Max time a connection can be open

	// Establish the connection pool
	emsPgxConnPool, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Fatal("unable to create connection pool: %w", err)
		os.Exit(1)
	}

	// Verify connectivity with a health check
	err = emsPgxConnPool.Ping(ctx)
	if err != nil {
		emsPgxConnPool.Close()
		log.Fatal("pool health check failed: %w", err)
		os.Exit(1)
	}
}

func GetEmsPgxConnPool() *pgxpool.Pool {
	CreateConnectionPool()
	return emsPgxConnPool
}
