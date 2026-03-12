package service

import (
	"fmt"
	"log"

	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"github.com/wwwmonster/eShopApp/go/v2/internal/dto"
	"github.com/wwwmonster/eShopApp/go/v2/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) Register(input *dto.UserRegister) (string, error) {
	log.Println(input)

	newUser, err := s.Repo.CreateUser(domain.User{
		Password: input.Password,
		Email:    input.Email,
		Phone:    input.Phone,
	})
	if err != nil {
		return "", nil
	}

	newUserInfo := fmt.Sprintf("%v", newUser.ID, newUser.Email, newUser.UserType)

	return newUserInfo, nil
}

func (s UserService) FindUserByEmail(input *dto.UserRegister, email string) (*domain.User, error) {
	fmt.Println("---ud.Email---", input.Email)

	fmt.Println("---------", input)
	user := domain.User{
		ID:        123,
		FirstName: "Alex",
		LastName:  "Li",
		Email:     "alexLi@gmail.com",
	}
	return &user, nil
}

func (s UserService) Login(input any) (string, error) {
	return "", nil
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
