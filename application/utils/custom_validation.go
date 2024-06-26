package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

type CustomValidationStruct struct {
	validate *validator.Validate
	ctx *fiber.Ctx
}

type CustomValidationInterface interface {
	// nama validatenya file_custom_validate
	// type validatenya *multipart.FileHeader
	FileCustomValidate()
}

func NewCustomValidation(validate *validator.Validate, ctx *fiber.Ctx) (CustomValidationInterface) {
	customValidationStruct := &CustomValidationStruct{
		validate: validate,
		ctx: ctx,
	}

	// register custom validation
	customValidationStruct.FileCustomValidate()
	customValidationStruct.ImageCustomValidate()
	return customValidationStruct
}

// nama validatenya file_custom_validate
// type validatenya *multipart.FileHeader
func (cl CustomValidationStruct) FileCustomValidate() {
	type fileValidation struct {
		File string `validate:"file"`
	}
	cl.validate.RegisterValidation("file_custom_validate", func (fl validator.FieldLevel) bool {
		// get type field
		typeField, _, _ := fl.ExtractType(fl.Field())
		// check jika type field bukan file atau bisa disebut multipart.FileHeader
		if typeField.Type().String() != "multipart.FileHeader" {
			return false
		}
		// get file from request
		file := fl.Field().Interface().(multipart.FileHeader)
		unixName := strconv.FormatInt(time.Now().UnixMicro(), 10)
		// set file name
		fileName := unixName + file.Filename
		pathTmp := GetApplicationPath() + "/storage/private/tmp/"
		// save file to tmp directory
		SaveFile(cl.ctx, &file, pathTmp, fileName)
		// delete file after validate
		defer os.Remove(pathTmp + fileName)

		// validate apakah ini file
		err := cl.validate.Struct(fileValidation{File: pathTmp + fileName})
		if err != nil {
			return false
		}
			return true
	})
}

// nama validatenya image_custom_validate
// type validatenya *multipart.FileHeader
func (cl CustomValidationStruct) ImageCustomValidate() {
	type fileValidation struct {
		File string `validate:"image"`
	}
	cl.validate.RegisterValidation("image_custom_validate", func (fl validator.FieldLevel) bool {
		// get type field
		typeField, _, _ := fl.ExtractType(fl.Field())
		// check jika type field bukan file atau bisa disebut multipart.FileHeader
		if typeField.Type().String() != "multipart.FileHeader" {
			return false
		}
		// get file from request
		file := fl.Field().Interface().(multipart.FileHeader)
		unixName := strconv.FormatInt(time.Now().UnixMicro(), 10)
		// set file name
		fileName := unixName + file.Filename
		pathTmp := GetApplicationPath() + "/storage/private/tmp/"
		// save file to tmp directory
		SaveFile(cl.ctx, &file, pathTmp, fileName)
		// delete file after validate
		defer os.Remove(pathTmp + fileName)

		// validate apakah ini file
		err := cl.validate.Struct(fileValidation{File: pathTmp + fileName})
		if err != nil {
			return false
		}
		return true
	})
}
