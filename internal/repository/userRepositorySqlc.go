package repository

import (
	"context"
	"log"

	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"github.com/wwwmonster/eShopApp/go/v2/internal/sqlc/eshopsqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepositorySqlc struct {
	connPool *pgxpool.Pool
	// ctx      *context.Context
	// conn     *pgx.Conn
}

// func NewUserRepositorySqlc(db *gorm.DB) UserRepository {
// 	return &userRepositorySqlc{db: db}
// }

func NewUserRepositorySqlc(cconnPool *pgxpool.Pool) UserRepository {
	return &userRepositorySqlc{connPool: cconnPool}
}

func (r userRepositorySqlc) CreateUser(usr domain.User) (domain.User, error) {
	// err := r.db.Create(&usr).Error
	// if err != nil {
	// 	log.Printf("create user error %v", err)
	// 	return domain.User{}, errors.New("failed to create user")
	// }

	return usr, nil
}

func (r userRepositorySqlc) FindUser(email string) (domain.User, error) {

	// ctx := context.Background()

	// conn, err := pgx.Connect(ctx, "host=127.0.0.1 user=root password=root dbname=online-shopping port=5432 sslmode=disable")

	// if err != nil {
	// 	log.Println("======userRepositorySqlc=======", err)
	// 	return domain.User{}, err
	// }

	ctx := context.Background()

	tx, _ := r.connPool.Begin(ctx)
	defer tx.Rollback(ctx)

	queries := eshopsqlc.New(r.connPool)

	// list all authors
	dbuser, err := queries.Getuser(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	log.Println("sqlc user.ID: ", dbuser.ID)
	user := domain.User{
		ID:        uint(dbuser.ID),
		FirstName: dbuser.FirstName.String,
		LastName:  dbuser.LastName.String,
		Email:     dbuser.Email,
	}
	return user, nil

	/*
		// create an author
		insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
			Name: "Brian Kernighan",
			Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
		})
		if err != nil {
			return err
		}
		log.Println(insertedAuthor)

		// get the author we just inserted
		fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
		if err != nil {
			return err
		}

		// prints true
		log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
		return nil

		var user domain.User

		err := r.db.Preload("Address").First(&user, "email=?", email).Error
		if err != nil {
			log.Printf("find user error %v", err)
			return domain.User{}, errors.New("user does not exist")
		}

		return user, nil
	*/
}

func (r userRepositorySqlc) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	// err := r.db.Preload("Address").
	// 	Preload("Cart").
	// 	Preload("Orders").
	// 	First(&user, id).Error
	// if err != nil {
	// 	log.Printf("find user error %v", err)
	// 	return domain.User{}, errors.New("user does not exist")
	// }

	return user, nil
}

func (r userRepositorySqlc) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	// err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error
	// if err != nil {
	// 	log.Printf("error on update %v", err)
	// 	return domain.User{}, errors.New("failed update user")
	// }

	return user, nil

}

func (r userRepositorySqlc) CreateBankAccount(e domain.BankAccount) error {
	// log.Println("CreateBankAccount...")
	// return r.db.Create(&e).Error
	return nil
}
