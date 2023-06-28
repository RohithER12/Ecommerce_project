package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "rohith")
	os.Setenv("DB_PASSWORD", "892328")
	os.Setenv("DB_NAME", "firstproject")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("SERVER_ADDRESS", "localhost:8080")
	os.Setenv("TWILIO_ACCOUNT_SID", "AC928c7b066593163befdf99e5f62df6b2")
	os.Setenv("TWILIO_AUTH_TOKEN", "81fb8e18e1224334660e5600a569b81a")
	os.Setenv("TWILIO_VERIFY_SERVICE_SID", "VAb7547e8b792265add4f519eb81d4de41")
	os.Setenv("RAZORPAY_KEY", "rzp_test_fLV8MhxNbQvHj2")
	os.Setenv("RAZORPAY_SECRET", "BApvy9EJwAaCeQT3ouZsNP5v")
	os.Setenv("JWT_SECRET", "testjwtsecret")

	cfg, err := LoadConfig()
	if err != nil {
		t.Errorf("error loading config: %v", err)
	}

	if cfg.DB.Host != "localhost" {
		t.Errorf("expected DB host: localhost, got: %s", cfg.DB.Host)
	}

	if cfg.DB.User != "rohith" {
		t.Errorf("expected DB user: rohith, got: %s", cfg.DB.User)
	}

	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("SERVER_ADDRESS")
	os.Unsetenv("TWILIO_ACCOUNT_SID")
	os.Unsetenv("TWILIO_AUTH_TOKEN")
	os.Unsetenv("TWILIO_VERIFY_SERVICE_SID")
	os.Unsetenv("RAZORPAY_KEY")
	os.Unsetenv("RAZORPAY_SECRET")
	os.Unsetenv("JWT_SECRET")
}
