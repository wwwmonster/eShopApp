package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest"
	"github.com/wwwmonster/eShopApp/go/v2/internal/dto"
	"github.com/wwwmonster/eShopApp/go/v2/internal/service"
)

type UserHandler struct {
	// service
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	fmt.Println("sur: ", rh)
	svc := service.UserService{}
	fmt.Printf("svc point address ---1---: %p\n", &svc)

	app := rh.App
	userHandler := UserHandler{svc: svc}
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
	fmt.Printf("svc point address ---2---: %p\n", &h.svc)

	user := new(dto.UserRegister)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse JSON")
	}

	fmt.Println(user.Email)
	token, err := h.svc.Register(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to login")
	}
	//	return ctx.Status(http.StatusOK).JSON(dbuser)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is register token: " + token,
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
