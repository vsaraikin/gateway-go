package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() (string, string, error) {
	if err := godotenv.Load("config/.env"); err != nil {
		return "", "", err
	}

	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	if apiKey == "" || secretKey == "" {

		return "", "", fmt.Errorf("API_KEY or SECRET_KEY not found in .env file")
	}

	return apiKey, secretKey, nil
}
