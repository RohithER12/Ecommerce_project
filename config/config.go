package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
)

type Config struct {
	DB struct {
		Host     string
		User     string
		Password string
		Name     string
		Port     string
		DSN      string
	}
	Server struct {
		Address string
	}
	JWT struct {
		Secret string
	}
	Twilio struct {
		AccountSID       string
		AuthToken        string
		VerifyServiceSID string
	}
	RazorPay struct {
		RazorpayKey    string
		RazorpaySecret string
	}
}

var config = Config{}
var client *twilio.RestClient

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	config := &Config{
		DB: struct {
			Host     string
			User     string
			Password string
			Name     string
			Port     string
			DSN      string
		}{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			DSN:      "",
		},
		Server: struct {
			Address string
		}{
			Address: os.Getenv("SERVER_ADDRESS"),
		},
		JWT: struct {
			Secret string
		}{
			Secret: "",
		},
		Twilio: struct {
			AccountSID       string
			AuthToken        string
			VerifyServiceSID string
		}{
			AccountSID:       os.Getenv("TWILIO_ACCOUNT_SID"),
			AuthToken:        os.Getenv("TWILIO_AUTH_TOKEN"),
			VerifyServiceSID: os.Getenv("TWILIO_VERIFY_SERVICE_SID"),
		},
		RazorPay: struct {
			RazorpayKey    string
			RazorpaySecret string
		}{
			RazorpayKey:    os.Getenv("RAZORPAY_KEY"),
			RazorpaySecret: os.Getenv("RAZORPAY_SECRET"),
		},
	}

	fmt.Println(
		"\nenv data\n", config.DB,
		"\n AccountSID\n", config.Twilio.AccountSID,
		"\n AuthToken\n", config.Twilio.AuthToken,
		"\n AuthToken\n", config.Server,
	)
	// Generate secret key or retrieve it from an environment variable
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey, err := generateSecretKey(32)
		if err != nil {
			return nil, fmt.Errorf("error generating secret key: %w", err)
		}
		config.JWT.Secret = secretKey
	} else {
		config.JWT.Secret = secretKey
	}

	config.DB.DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB.Host, config.DB.User, config.DB.Password, config.DB.Name, config.DB.Port)

	return config, nil
}

func GetJWTConfig() string {
	return config.JWT.Secret
}

func init() {
	// config.Twilio.AccountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	// config.Twilio.AuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	// config.Twilio.VerifyServiceSID = os.Getenv("TWILIO_VERIFY_SERVICE_SID")

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.Twilio.AccountSID,
		Password: config.Twilio.AuthToken,
	})
}

func GetTwilioClient() *twilio.RestClient {
	return client
}

func GetJWTSecret() string {
	return config.JWT.Secret
}

func generateSecretKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	secretKey := base64.URLEncoding.EncodeToString(randomBytes)
	return secretKey, nil
}
