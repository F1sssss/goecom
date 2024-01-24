package utils

import (
	"os"
)

var (
	// Email
	EMAIL_USERNAME string
	EMAIL_PASSWORD string
	EMAIL_HOST     string
	EMAIL_PORT     int
	SERVER_URL     string
)

func InitConfig() error {

	EMAIL_USERNAME = os.Getenv("EMAIL_HOST_USER")
	EMAIL_PASSWORD = os.Getenv("EMAIL_HOST_PASSWORD")
	EMAIL_HOST = os.Getenv("EMAIL_HOST")
	SERVER_URL = os.Getenv("SERVER_URL")

	return nil
}
