package helpers

import (
	"70_Off/infrastructure/config"
	"fmt"

	"github.com/razorpay/razorpay-go"
)

func GenerateRazorpayOrder(razorPayAmount uint) (razorpayOrderID string, err error) {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration:", err)
		return
	}

	razorpayKey := config.RazorPay.RazorpayKey
	razorpaySecret := config.RazorPay.RazorpaySecret

	client := razorpay.NewClient(razorpayKey, razorpaySecret)

	data := map[string]interface{}{
		"amount":   razorPayAmount,
		"currency": "INR",
		"receipt":  "101",
	}

	razorpayRes, err := client.Order.Create(data, nil)
	if err != nil {
		return razorpayOrderID, fmt.Errorf("failed to create razorpay order for amount %v", razorPayAmount)
	}

	razorpayOrderID = razorpayRes["id"].(string)

	return razorpayOrderID, nil
}
