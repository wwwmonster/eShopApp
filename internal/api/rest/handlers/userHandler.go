package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest"
)

type UserHandler struct {
	// service
}

func SetupUserRoutes(rh *rest.RestHandler) {
	fmt.Println("sur: ", rh)
	app := rh.App
	userHandler := UserHandler{}
	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	app.Get("/verify", userHandler.GetVerificationCode)
	app.Post("/verify", userHandler.Verify)
	app.Get("/profile", userHandler.GetProfile)
	app.Post("/profile", userHandler.CreateProfile)

	app.Post("/cart", userHandler.AddToCart)
	app.Get("/cart", userHandler.GetCart)
	app.Get("/order", userHandler.Register)
	app.Get("/order/:id", userHandler.Register)

	app.Post("/become-seller", userHandler.BecomeSeller)
}

type UserData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	ud := new(UserData)
	if err := ctx.BodyParser(ud); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to parse JSON")
	}

	fmt.Println(ud)
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is register...",
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is Login...",
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is GetVerificationCode...",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is Verify...",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is GetProfile...",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is CreateProfile...",
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is AddToCart...",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is GetCart...",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is GetOrders...",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is GetOrder...",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is BecomeSeller...",
	})
}
