package helper2

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(key string, defau string) string {
	loadEnv()
	data := os.Getenv(key)

	if data == "" {
		data = defau
	}
	return data
}

func loadEnv() {
	rootPath := GetRootPath()

	err := godotenv.Load(rootPath + `/.env`)

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}




