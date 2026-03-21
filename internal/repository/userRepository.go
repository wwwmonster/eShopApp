package repository

import (
	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetDb() *gorm.DB
	CreateUser(usr domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
	CreateBankAccount(e domain.BankAccount) error

	BecomeBuyer(u *domain.User, e *domain.BankAccount) error
	WithTx(tx *gorm.DB) UserRepository

	/*
		// Cart
		FindCartItems(uId uint) ([]domain.Cart, error)
		FindCartItem(uId uint, pId uint) (domain.Cart, error)
		CreateCart(c domain.Cart) error
		UpdateCart(c domain.Cart) error
		DeleteCartById(id uint) error
		DeleteCartItems(uId uint) error

		// Order
		CreateOrder(o domain.Order) error
		FindOrders(uId uint) ([]domain.Order, error)
		FindOrderById(id uint, uId uint) (domain.Order, error)

		// Profile
		CreateProfile(e domain.Address) error
		UpdateProfile(e domain.Address) error
	*/
}
