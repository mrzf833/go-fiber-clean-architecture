package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

const projectDirName = "go-fiber-clean-architecture" // change to relevant project name

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func Config(key string, defau string) string {
	// load .env file
	loadEnv()

	data := os.Getenv(key)

	if data == "" {
		data = defau
	}
	return data
}




