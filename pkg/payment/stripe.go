package payment

import (
	"errors"
	"fmt"
	"log"

	"github.com/stripe/stripe-go/v85"
	"github.com/stripe/stripe-go/v85/paymentintent"
)

type PaymentClient interface {
	CreatePayment(amount float64, userId uint, orderId string) (*stripe.PaymentIntent, error)
	GetPaymentStatus(pId string) (*stripe.PaymentIntent, error)
}

type payment struct {
	stripeSecretKey string
}

func NewPaymentClient(stripeSecretKey string) PaymentClient {
	return &payment{
		stripeSecretKey: stripeSecretKey,
	}
}

func (p payment) CreatePayment(amount float64, userId uint, orderId string) (*stripe.PaymentIntent, error) {
	stripe.Key = p.stripeSecretKey

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount * 100)), // Amount in cents
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Metadata: map[string]string{
			"user_id":  fmt.Sprintf("%d", userId),
			"order_id": orderId,
		},
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		log.Printf("Error creating payment intent: %v", err)
		return nil, errors.New("create payment intent failed")
	}
	return pi, nil
}

func (p payment) GetPaymentStatus(pId string) (*stripe.PaymentIntent, error) {
	stripe.Key = p.stripeSecretKey
	params := &stripe.PaymentIntentParams{}
	result, err := paymentintent.Get(pId, params)

	if err != nil {
		log.Printf("Error getting payment intent: %v", err)
		return nil, errors.New("get payment intent failed")
	}
	return result, nil
}
