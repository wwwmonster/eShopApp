package helper

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wwwmonster/eShopApp/go/v2/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret         string
	PasswordMinLen int
}

func SetupAuth(s string, pml int) Auth {
	return Auth{Secret: s, PasswordMinLen: pml}
}

func (a Auth) CreateHashedPassword(p string) (string, error) {
	if len(p) < a.PasswordMinLen {
		return "", errors.New("Password length should be at least 6 charactors... ")
	}
	hashedPasswrod, err := bcrypt.GenerateFromPassword([]byte(p), 10)

	if err != nil {
		return "", err
	}
	return string(hashedPasswrod), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("Required inputs are missing to generate token... ")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		return "", errors.New("Unable to sign the token ")
	}

	return tokenStr, nil
}

func (a Auth) VerifyPassword(pP string, hP string) error {
	if len(pP) < 6 {
		return errors.New("Password length should be at least 6 charactors... ")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP)); err != nil {
		return errors.New("Invalid Password")
	}
	return nil
}

func (a Auth) VerifyToken(t string) (domain.User, error) {

	log.Println(t)
	// if strings.HasPrefix(t, "Bearer ") {
	// 	return domain.User{}, errors.New("Invalid JWT token...no 'Bearer ' ")
	// }

	tokenArr := strings.Split(t, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, errors.New("Invalid JWT token...cannot split ")
	}

	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("Invalid JWT token... token 0 is not Bearer")
	}

	token, err := jwt.Parse(tokenArr[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)
		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]

	if len(authHeader) < 1 {
		return ctx.Status((401)).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  "no JWT",
		})
	}
	user, err := a.VerifyToken(authHeader[0])
	log.Println(err)
	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status((401)).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	}
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)
}

func (a Auth) GenerateCode() (int, error) {
	return GenerateRandomNumbers(6)
}

func (a Auth) AuthorizeSeller(ctx *fiber.Ctx) error {

	authHeader := ctx.GetReqHeaders()["Authorization"]
	user, err := a.VerifyToken(authHeader[0])

	if err != nil {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	} else if user.ID > 0 && user.UserType == domain.SELLER {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  errors.New("please join seller program to manage products"),
		})
	}

}
