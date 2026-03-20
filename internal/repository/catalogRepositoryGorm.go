package repository

import (
	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"gorm.io/gorm"
)

type catalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{db: db}
}

func (c catalogRepository) CreateUser(usr domain.User) error {
	return nil
}
