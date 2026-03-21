package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api/rest"
	"github.com/wwwmonster/eShopApp/go/v2/internal/dto"
	"github.com/wwwmonster/eShopApp/go/v2/internal/repository"
	"github.com/wwwmonster/eShopApp/go/v2/internal/service"
)

type CatalogHandler struct {
	// service
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	svc := service.CatalogService{Repo: repository.NewCatalogRepository(rh.Db), Auth: rh.Auth, Config: rh.Config}
	// svc1 := service.CatalogService{Repo: repository.NewUserRepositorySqlc(rh.ConnPool), Auth: rh.Auth}
	fmt.Printf("svc point address ---1---: %p\n", &svc)
	// fmt.Printf("svc1 point address ---1---: %p\n", &svc1)

	app := rh.App
	handler := CatalogHandler{svc: svc}

	// public
	// listing products and categories
	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProduct)
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryById)

	// private
	// manage products and categories
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)
	// Categories
	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	// Products
	selRoutes.Post("/products", handler.CreateProducts)
	selRoutes.Get("/products", handler.GetProducts)
	selRoutes.Get("/products/:id", handler.GetProduct)
	selRoutes.Put("/products/:id", handler.EditProduct)
	selRoutes.Patch("/products/:id", handler.UpdateStock) // update stock
	selRoutes.Delete("/products/:id", handler.DeleteProduct)

}
func (h CatalogHandler) GetProducts(ctx *fiber.Ctx) error {

	// products, err := h.svc.GetProducts()
	// if err != nil {
	// 	return rest.ErrorMessage(ctx, 404, err)
	// }

	return rest.SuccessResponse(ctx, "products", nil)
}

func (h CatalogHandler) GetProduct(ctx *fiber.Ctx) error {

	// id, _ := strconv.Atoi(ctx.Params("id"))

	// product, err := h.svc.GetProductById(id)
	// if err != nil {
	// 	return rest.BadRequestError(ctx, "product not found")
	// }

	return rest.SuccessResponse(ctx, "product", nil)
}

func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {
	if cats, err := h.svc.GetCategories(); err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	} else {
		return rest.SuccessResponse(ctx, "categories", cats)
	}
}
func (h CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	log.Println(id)

	if cat, err := h.svc.GetCategory(id); err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	} else {
		return rest.SuccessResponse(ctx, "category", cat)
	}
}

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	log.Println("user: ", user.Email)

	req := new(dto.CreateCategoryRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestError(ctx, "create category request is not valid")
	}

	if err := h.svc.CreateCategory(req); err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "category created successfully", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	req := new(dto.CreateCategoryRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestError(ctx, "update category request is not valid")
	}

	if updatedCat, err := h.svc.EditCategory(id, req); err != nil {
		return rest.InternalError(ctx, err)
	} else {
		return rest.SuccessResponse(ctx, "edit category", updatedCat)
	}
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	if err := h.svc.DeleteCategory(id); err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.SuccessResponse(ctx, "category deleted successfully", nil)
}

func (h CatalogHandler) CreateProducts(ctx *fiber.Ctx) error {

	req := new(dto.CreateProductRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestError(ctx, "create product request is not valid")
	}

	user := h.svc.Auth.GetCurrentUser(ctx)
	if err := h.svc.CreateProduct(req, user); err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "product created successfully", nil)
}

func (h CatalogHandler) EditProduct(ctx *fiber.Ctx) error {

	// id, _ := strconv.Atoi(ctx.Params("id"))
	// req := dto.CreateProductRequest{}
	// err := ctx.BodyParser(&req)
	// if err != nil {
	// 	return rest.BadRequestError(ctx, "edit product request is not valid")
	// }
	// user := h.svc.Auth.GetCurrentUser(ctx)
	// product, err := h.svc.EditProduct(id, req, user)
	// if err != nil {
	// 	return rest.InternalError(ctx, err)
	// }
	// return rest.SuccessResponse(ctx, "edit product", product)
	return rest.SuccessResponse(ctx, "edit product", nil)
}

func (h CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	// id, _ := strconv.Atoi(ctx.Params("id"))
	// req := dto.UpdateStockRequest{}
	// err := ctx.BodyParser(&req)
	// if err != nil {
	// 	return rest.BadRequestError(ctx, "update stock request is not valid")
	// }
	// user := h.svc.Auth.GetCurrentUser(ctx)

	// product := domain.Product{
	// 	ID:     uint(id),
	// 	Stock:  uint(req.Stock),
	// 	UserId: int(user.ID),
	// }

	// updatedProduct, err := h.svc.UpdateProductStock(product)

	// return rest.SuccessResponse(ctx, "update stock ", updatedProduct)
	return rest.SuccessResponse(ctx, "update stock ", nil)
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {

	// id, _ := strconv.Atoi(ctx.Params("id"))
	// // need to provide user id to verify ownership
	// user := h.svc.Auth.GetCurrentUser(ctx)
	// err := h.svc.DeleteProduct(id, user)

	// return rest.SuccessResponse(ctx, "Delete product ", err)
	return rest.SuccessResponse(ctx, "Delete product ", nil)
}
