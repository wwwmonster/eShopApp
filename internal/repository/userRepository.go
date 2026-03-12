package repository

import (
	"errors"
	"log"

	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(usr domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
	CreateBankAccount(e domain.BankAccount) error
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

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {
	err := r.db.Create(&usr).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return usr, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User

	err := r.db.Preload("Address").First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.Preload("Address").
		Preload("Cart").
		Preload("Orders").
		First(&user, id).Error
	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error
	if err != nil {
		log.Printf("error on update %v", err)
		return domain.User{}, errors.New("failed update user")
	}

	return user, nil
}

func (r userRepository) CreateBankAccount(e domain.BankAccount) error {
	log.Println("CreateBankAccount...")
	return r.db.Create(&e).Error
}
