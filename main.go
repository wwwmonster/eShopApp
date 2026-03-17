package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wwwmonster/eShopApp/go/v2/configs"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api"
)

func main() {
	fmt.Println("start main...")
	log.Println("============")
	config, err := configs.SetupEnv()
	if err != nil {
		os.Exit(0)
	}
	api.StartServer(config)
	//	testing.Testing()
}

func main2() {
	if os.Getenv("APP_ENV") == "dev" {
		fmt.Println("11111")
	} else {
		fmt.Println("22222")
	}

}
