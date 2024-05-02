package helper

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
)

func SaveFile(c *fiber.Ctx, fileheader *multipart.FileHeader, pathDirectory string, nameFile string) error  {
	err := MakeDirectoryIfNotExists(pathDirectory)
	if err != nil {
		return err
	}

	return c.SaveFile(fileheader, pathDirectory + "/" + nameFile)
}

func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModeDir|0755)
	}
	return nil
}
