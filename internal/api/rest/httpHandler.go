package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
	"github.com/wwwmonster/eShopApp/go/v2/internal/helper"
	"github.com/wwwmonster/eShopApp/go/v2/pkg/payment"
	"gorm.io/gorm"
)

type RestHandler struct {
	App *fiber.App
	Db  *gorm.DB
	// Ctx      *context.Context
	// Conn     *pgx.Conn
	ConnPool      *pgxpool.Pool
	Auth          helper.Auth
	Config        configs.AppConfig
	PaymentClient payment.PaymentClient
}
