package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {
	//	if os.Getenv("APP_ENV") == "dev" {
	if enverr := godotenv.Load(); enverr != nil {
		return AppConfig{}, errors.New("env loading failed")
	}
	//	}
	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env loading HTTP_PORT")
	}
	cfg = AppConfig{
		ServerPort: httpPort,
	}
	return cfg, nil
}
