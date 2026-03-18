package service

import (
	"errors"
	"log"
	"strconv"
	"time"

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

func (s UserService) isVerifiedUser(id uint) bool {
	currentUder, err := s.Repo.FindUserById(id)
	return err == nil && currentUder.Verified
}

func (s UserService) GetVerificationCode(u domain.User) (int, error) {
	if s.isVerifiedUser(u.ID) {
		return 0, errors.New("user already verified")
	}

	code, err := s.Auth.GenerateCode()

	if err != nil {
		return 0, err
	}
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   strconv.Itoa(code),
	}

	_, err = s.Repo.UpdateUser(u.ID, user)
	if err != nil {
		return 0, errors.New("unable to update verification code")
	}

	return code, nil
}

func (s UserService) VerifyCode(id uint, code string) error {
	if s.isVerifiedUser(id) {
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	updatedUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, updatedUser)
	if err != nil {
		return errors.New("unable to update verification code")
	}

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
