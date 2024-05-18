package helper2

import (
	"go-fiber-clean-architecture/application/config/config2"
	"os"
	"regexp"
)

func GetRootPath() string {
	projectName := regexp.MustCompile(`^(.*` + config2.ProjectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}
