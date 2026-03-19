package notification

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
)

type NotificationClient interface {
	SendSMS(phone string, msg string) error
}

type notificationClient struct {
	config configs.AppConfig
}

func (c notificationClient) SendSMS(phone string, msg string) error {
	// accountSid := "AC3ecb8f334aa7ba7cf348914f4ad906fa"
	// authToken := "7587c323f1aeee77ef41871b01d61cea"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.config.AccountSid,
		Password: c.config.AuthToken,
	})

	// params := &openapi.CreateMessageParams{}
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(c.config.FromPhone)
	params.SetBody(msg)

	if c.config.IsSendSMS {
		log.Println("sending SMS verification code to user's phone")
		resp, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println("Error sending SMS message: " + err.Error())
		} else {
			response, _ := json.Marshal(*resp)
			fmt.Println("Response: " + string(response))
		}
	} else {
		log.Println("skip sending SMS")
	}

	return nil
}

func NewNotificationClient(config configs.AppConfig) NotificationClient {
	return &notificationClient{config: config}
}

func NewNotificationClient1() NotificationClient {
	return &notificationClient{}
}

func SendSMSTest(phone string, msg string) error {
	// accountSid := "AC3ecb8f334aa7ba7cf348914f4ad906fa"
	// authToken := "7587c323f1aeee77ef41871b01d61cea"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: "AC3ecb8f334aa7ba7cf348914f4ad906fa",
		Password: "7587c323f1aeee77ef41871b01d61cea",
	})

	// params := &openapi.CreateMessageParams{}
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom("+18734578151")
	params.SetBody(msg)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}
