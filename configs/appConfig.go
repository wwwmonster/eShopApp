package configs

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	DbConnType string
	AppSecret  string
	AccountSid string
	AuthToken  string
	FromPhone  string
	IsSendSMS  bool
}

const (
	GORM = iota
	SQLC
)

func SetupEnv() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		if enverr := godotenv.Load(); enverr != nil {
			return AppConfig{}, errors.New("env loading failed")
		}
	} else {
		if enverr := godotenv.Load(); enverr != nil {
			return AppConfig{}, errors.New("env loading failed")
		}
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env loading HTTP_PORT failed")
	}
	Dsn := os.Getenv("DSN")

	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("env loading DSN failed")
	}
	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env loading APP_SECRET failed")
	}

	accountSid := os.Getenv("ACCOUNT_SID")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env loading ACCOUNT_SID failed")
	}

	authToken := os.Getenv("AUTH_TOKEN") + os.Getenv("AUTH_TOKEN_2")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env loading AUTH_TOKEN failed")
	}

	fromPhone := os.Getenv("FROM_PHONE")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env loading FROM_PHONE failed")
	}

	isSendSMS, err := strconv.ParseBool(os.Getenv("IS_SEND_NOTIFIACTION_CODE"))
	if len(appSecret) < 1 || err != nil {
		return AppConfig{}, errors.New("env loading IS_SEND_NOTIFIACTION_CODE failed")
	}

	cfg = AppConfig{
		ServerPort: httpPort,
		Dsn:        Dsn,
		AppSecret:  appSecret,
		AccountSid: accountSid,
		AuthToken:  authToken,
		FromPhone:  fromPhone,
		IsSendSMS:  isSendSMS,
	}
	return cfg, nil
}
