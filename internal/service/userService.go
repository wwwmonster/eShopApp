package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/wwwmonster/eShopApp/go/v2/configs"
	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"github.com/wwwmonster/eShopApp/go/v2/internal/dto"
	"github.com/wwwmonster/eShopApp/go/v2/internal/helper"
	"github.com/wwwmonster/eShopApp/go/v2/internal/repository"
	"github.com/wwwmonster/eShopApp/go/v2/pkg/notification"
	"gorm.io/gorm"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config configs.AppConfig
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

func (s UserService) GetVerificationCode(u domain.User) (string, error) {
	if s.isVerifiedUser(u.ID) {
		return "", errors.New("user already verified")
	}

	code, err := s.Auth.GenerateCode()

	if err != nil {
		return "", err
	}
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   strconv.Itoa(code),
	}

	_, err = s.Repo.UpdateUser(u.ID, user)
	if err != nil {
		return "", errors.New("unable to update verification code")
	}

	//send SMS
	user, _ = s.Repo.FindUserById(u.ID)
	log.Println("user.phone: ", user.Phone)
	notificationClient := notification.NewNotificationClient(s.Config)
	msg := fmt.Sprintf("Your verification code is %v", code)
	if err = notificationClient.SendSMS(user.Phone, msg); err != nil {
		return "", errors.New("Failed to send SMS")
	}

	return strconv.Itoa(code), nil
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

func (s UserService) BecomeBuyer1(id uint, input dto.SellerInput) (string, error) {
	user, _ := s.Repo.FindUserById(id)
	if user.UserType == domain.SELLER {
		// return "", errors.New("you have already joined seller program")
	}
	log.Println("-------BecomeBuyer-----1-----")
	s.Repo.BecomeBuyer(&domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.PhoneNumber,
		UserType:  domain.SELLER,
	},
		&domain.BankAccount{
			BankAccount: input.BankAccountNumber,
			SwiftCode:   input.SwiftCode,
			PaymentType: input.PaymentType,
			UserId:      id,
		},
	)
	return s.Auth.GenerateToken(user.ID, user.Email, domain.SELLER)
}

func (s UserService) BecomeBuyer2(id uint, input dto.SellerInput) (string, error) {
	log.Println("-------BecomeBuyer-----2-----")
	user, _ := s.Repo.FindUserById(id)
	if user.UserType == domain.SELLER {
		// return "", errors.New("you have already joined seller program")
	}

	db := s.Repo.GetDb()
	log.Printf("============service db %p/", db)
	if err := db.Transaction(func(tx *gorm.DB) error {
		txRepo := repository.NewUserRepository(tx) // 👈 inject tx here
		// txRepo := s.Repo.WithTx(tx) // 👈 inject tx here
		if _, err := txRepo.UpdateUser(id, domain.User{
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Phone:     input.PhoneNumber,
			UserType:  domain.SELLER,
		}); err != nil {
			return errors.New("Faile to update user to seller")
		} else {
			if err := txRepo.CreateBankAccount(
				domain.BankAccount{
					BankAccount: input.BankAccountNumber,
					SwiftCode:   input.SwiftCode,
					PaymentType: input.PaymentType,
					UserId:      id,
				}); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return "", errors.New("Faile to update user to seller")
	}
	return s.Auth.GenerateToken(user.ID, user.Email, domain.SELLER)
}

func (s UserService) BecomeBuyer(id uint, input dto.SellerInput) (string, error) {

	user, _ := s.Repo.FindUserById(id)
	if user.UserType == domain.SELLER {
		// return "", errors.New("you have already joined seller program")
	}

	if user, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.PhoneNumber,
		UserType:  domain.SELLER,
	}); err != nil {
		return "", errors.New("Faile to update user to seller")
	} else {
		if err := s.Repo.CreateBankAccount(
			domain.BankAccount{
				BankAccount: input.BankAccountNumber,
				SwiftCode:   input.SwiftCode,
				PaymentType: input.PaymentType,
				UserId:      id,
			}); err != nil {
			return "", err
		}
		return s.Auth.GenerateToken(user.ID, user.Email, domain.SELLER)
	}

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
