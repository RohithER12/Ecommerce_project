package helpers

import (
	"errors"
	"fmt"

	"70_Off/infrastructure/config"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	TWILIO_ACCOUNT_SID string
	TWILIO_AUTH_TOKEN  string
	TWILIO_SERVICE_ID  string
	client             *twilio.RestClient
)

func init() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration:", err)
		return
	}

	TWILIO_ACCOUNT_SID = config.Twilio.AccountSID
	TWILIO_AUTH_TOKEN = config.Twilio.AuthToken
	TWILIO_SERVICE_ID = config.Twilio.VerifyServiceSID

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})
}

func SendOtp(phone string) (string, error) {
	fmt.Println("\n", phone)
	to := "+91" + phone
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(TWILIO_SERVICE_ID, params)
	if err != nil {
		fmt.Println("Invalid PhoneNumber\n", err.Error())
		return "", err
	}

	var key string
	key = *resp.Sid
	fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	return key, nil
}

func CheckOtp(phoneNumber string, code string) error {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(TWILIO_SERVICE_ID, params)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Invalid otp")
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
		return nil
	} else {
		fmt.Println("Incorrect!")
		return errors.New("Invalid otp")
	}
}
