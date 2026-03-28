package repository

import (
	"errors"
	"log"

	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) WithTx(tx *gorm.DB) UserRepository {
	return userRepository{db: tx}
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

	err := r.db.Preload("BankAccount").Preload("Address").First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.Preload("Address").
		Preload("BankAccount").
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
	log.Printf("============UpdateUser db %p/", r.db)
	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error
	if err != nil {
		log.Printf("error on update %v", err)
		return domain.User{}, errors.New("failed update user")
	}

	return user, nil
}

func (r userRepository) CreateBankAccount(e domain.BankAccount) error {
	log.Printf("============CreateBankAccount db %p/", r.db)
	log.Println("CreateBankAccount...")
	return r.db.Create(&e).Error
}

func (r userRepository) GetDb() *gorm.DB {
	return r.db
}

func (r userRepository) BecomeBuyer(u *domain.User, e *domain.BankAccount) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if _, err := r.UpdateUser(u.ID, *u); err != nil {
			return errors.New("Faile to update user to seller")
		} else {
			if err := r.CreateBankAccount(*e); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return errors.New("Faile to update user to seller")
	}
	return nil
}

func (r userRepository) FindCartItem(uId uint, pId uint) (domain.Cart, error) {
	cartItem := domain.Cart{}
	err := r.db.Where("user_id=? AND product_id=?", uId, pId).First(&cartItem).Error
	return cartItem, err
}

// CreateCart implements [UserRepository].
func (r userRepository) CreateCart(c domain.Cart) error {
	return r.db.Create(&c).Error
}

// DeleteCartById implements [UserRepository].
func (r userRepository) DeleteCartById(id uint) error {
	return r.db.Delete(&domain.Cart{}, id).Error
}

// DeleteCartItems implements [UserRepository].
func (r userRepository) DeleteCartItems(uId uint) error {
	return r.db.Where("user_id=", uId).Delete(&domain.Cart{}).Error
}

// UpdateCart implements [UserRep`ository].
func (r userRepository) UpdateCart(c domain.Cart) error {

	// return r.db.Save(&c).Error
	var cart domain.Cart
	return r.db.Model(&cart).Clauses(clause.Returning{}).Where("id", c.ID).Updates(c).Error
}

func (r userRepository) FindCartItems(uId uint) ([]domain.Cart, error) {
	var carts []domain.Cart
	err := r.db.Where("user_id=?", uId).Find(&carts).Error
	return carts, err
}

func (r userRepository) CreateProfile(e domain.Address) error {
	if err := r.db.Create(&e).Error; err != nil {
		log.Printf("error on creating profile with address %v", err)
		return errors.New("failed to create profile")
	} else {
		return nil
	}
}

func (r userRepository) UpdateProfile(e domain.Address) error {
	err := r.db.Where("user_id=?", e.UserId).Updates(e).Error
	if err != nil {
		log.Printf("error on update profile with address %v", err)
		return errors.New("failed to create profile")
	}
	return nil

}
