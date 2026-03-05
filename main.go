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
