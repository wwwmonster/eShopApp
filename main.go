package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stripe/stripe-go/v85"
	"github.com/stripe/stripe-go/v85/price"
	"github.com/stripe/stripe-go/v85/product"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
	"github.com/wwwmonster/eShopApp/go/v2/internal/api"
	"github.com/wwwmonster/eShopApp/go/v2/pkg/notification"
)

func main() {
	StripeTest()
	fmt.Println("start main...")
	log.Println("============")
	config, err := configs.SetupEnv()
	if err != nil {
		os.Exit(0)
	}
	api.StartServer(config)
	//	testing.Testing()
}

func main1() {
	notification.SendSMSTest("6477126661", "Hello, this eSHop")

}

func StripeTest() {
	stripe.Key = "sk_test_51TFjpQGrfeXvFIUCDIpvPq7Wz7GF15N9J7sCg90pEKPuSBZjOY7ItW75zsZIcdeNdgJNWDzG42VwHBZzcvXICDEv00OirK0Sq6"

	product_params := &stripe.ProductParams{
		Name:        stripe.String("Starter Subscription"),
		Description: stripe.String("$12/Month subscription"),
	}
	starter_product, _ := product.New(product_params)

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Product:  stripe.String(starter_product.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(1200),
	}
	starter_price, _ := price.New(price_params)

	fmt.Println("Success! Here is your starter subscription product id: " + starter_product.ID)
	fmt.Println("Success! Here is your starter subscription price id: " + starter_price.ID)
}
