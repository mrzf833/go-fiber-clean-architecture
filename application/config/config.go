package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
