package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest"
	"github.com/wwwmonster/eShopApp/go/v2/internal/dto"
	"github.com/wwwmonster/eShopApp/go/v2/internal/repository"
	"github.com/wwwmonster/eShopApp/go/v2/internal/service"
)

type UserHandler struct {
	// service
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	svc := service.UserService{Repo: repository.NewUserRepository(rh.Db), Auth: rh.Auth, Config: rh.Config}
	// svc1 := service.UserService{Repo: repository.NewUserRepositorySqlc(rh.ConnPool), Auth: rh.Auth}
	fmt.Printf("svc point address ---1---: %p\n", &svc)
	// fmt.Printf("svc1 point address ---1---: %p\n", &svc1)

	app := rh.App
	userHandler := UserHandler{svc: svc}
	pubRoutes := app.Group("/users")
	pubRoutes.Post("/register", userHandler.Register)
	pubRoutes.Post("/login", userHandler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Get("/verify", userHandler.GetVerificationCode)
	pvtRoutes.Post("/verify", userHandler.Verify)
	pvtRoutes.Get("/profile", userHandler.GetProfile)
	pvtRoutes.Post("/profile", userHandler.CreateProfile)

	pvtRoutes.Post("/cart", userHandler.AddToCart)
	pvtRoutes.Get("/cart", userHandler.GetCart)
	pvtRoutes.Get("/order", userHandler.Register)
	pvtRoutes.Get("/order/:id", userHandler.Register)

	pvtRoutes.Post("/become-seller", userHandler.BecomeSeller)
}

type UserData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := new(dto.UserRegister)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse JSON")
	}

	token, err := h.svc.Register(user)
	if err != nil {
		log.Panic(err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to login")
	}
	//	return ctx.Status(http.StatusOK).JSON(dbuser)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is register",
		"token":   token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginUser := new(dto.UserRegister)
	if err := ctx.BodyParser(loginUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse JSON")
	}

	token, err := h.svc.Login(loginUser.Email, loginUser.Password)
	if err != nil {

		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Failed to login...",
			"error":   "UserName or Password not correct",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is Login...",
		"token":   token,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	code, err := h.svc.GetVerificationCode(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to GetVerificationCode...",
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "this is GetVerificationCode..." + code,
		// "code":    code,
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	codeStruct := new(dto.VerificationCodeInput)
	if err := ctx.BodyParser(codeStruct); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse JSON")
	}
	err := h.svc.VerifyCode(user.ID, codeStruct.Code)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Verified...",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	// user := new(dto.UserRegister)
	// if err := ctx.BodyParser(user); err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse JSON")
	// }

	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Println("inputEmail： ", user.Email)
	// dbuser, err := h.svc.FindUserByEmail("alex1@example.com")
	dbuser, err := h.svc.FindUserByEmail(user.Email)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to login")
	}
	return ctx.Status(http.StatusOK).JSON(dbuser)

	// return ctx.Status(http.StatusOK).JSON(&fiber.Map{
	// 	"message": "this is register token: " + "token",
	// })

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

	req := new(dto.SellerInput)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse Seller Input JSON")
	}

	user := h.svc.Auth.GetCurrentUser(ctx)
	if token, err := h.svc.BecomeBuyer(user.ID, *req); err != nil {
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Failed to BecomeSeller ...",
		})
	} else {
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "BecomeSeller Successfully...",
			"teken":   token,
		})
	}

}
