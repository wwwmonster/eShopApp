package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	DbConnType string
	AppSecret  string
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

	cfg = AppConfig{
		ServerPort: httpPort,
		Dsn:        Dsn,
		AppSecret:  appSecret,
	}
	return cfg, nil
}
