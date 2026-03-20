package service

import (
	"github.com/wwwmonster/eShopApp/go/v2/configs"
	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"github.com/wwwmonster/eShopApp/go/v2/internal/helper"
	"github.com/wwwmonster/eShopApp/go/v2/internal/repository"
)

type CatalogService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config configs.AppConfig
}

func (c CatalogService) GetProducts() ([]*domain.Product, error) {
	return []*domain.Product{}, nil
}

func (c CatalogService) GetProductById(id int) (*domain.Product, error) {
	return &domain.Product{}, nil
}
