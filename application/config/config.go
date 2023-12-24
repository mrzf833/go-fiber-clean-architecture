package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string, defau string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	data := os.Getenv(key)

	if data == "" {
		data = defau
	}
	return data
}




