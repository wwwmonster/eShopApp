package service

import (
	"errors"
	"log"

	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"github.com/wwwmonster/eShopApp/go/v2/internal/dto"
	"github.com/wwwmonster/eShopApp/go/v2/internal/helper"
	"github.com/wwwmonster/eShopApp/go/v2/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) Register(input *dto.UserRegister) (string, error) {
	log.Println(input)
	hPassword, err := s.Auth.CreateHashedPassword(input.Password)

	if err != nil {
		return "", nil
	}

	newUser, err := s.Repo.CreateUser(domain.User{
		Password: hPassword,
		Email:    input.Email,
		Phone:    input.Phone,
	})

	return s.Auth.GenerateToken(newUser.ID, newUser.Email, newUser.UserType)
}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {

	// sqlcUserS := repository.NewUserRepositorySqlc()
	// sqlcUser, _ := sqlcUserS.FindUser(email)

	// log.Println(sqlcUser)

	user, err := s.Repo.FindUser(email)
	if err != nil {
		return &domain.User{}, err
	}

	// user := domain.User{
	// 	ID:        123,
	// 	FirstName: "Alex",
	// 	LastName:  "Li",
	// 	Email:     "alexLi@gmail.com",
	// }
	return &user, nil
}

func (s UserService) Login(email string, password string) (string, error) {
	dbuser, err := s.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist")
	}

	err = s.Auth.VerifyPassword(password, dbuser.Password)

	if err != nil {
		return "", err
	}
	return s.Auth.GenerateToken(dbuser.ID, dbuser.Email, dbuser.UserType)
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProdile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) BecomeBuyer(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]domain.Cart, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]domain.Cart, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]domain.Order, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uid uint) (*domain.Order, error) {
	return nil, nil
}
