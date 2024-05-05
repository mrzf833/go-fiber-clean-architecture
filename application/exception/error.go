package exception

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/config"
)


type ErrorResponse struct {
	Field string
	Tag         string
	Value       interface{}
	Message 	interface{}
}

var (
	// error pesan untuk status 500
	ErrInternalServerError = errors.New("Internal Server Error")
	// error pesan untuk status 404
	ErrNotFound = errors.New("Your requested Item is not found")
	// error pesan untuk status 400
	ErrBadParamInput       = errors.New("Given Param is not valid")
)

// HandleError untuk menangani error
func HandleError(c *fiber.Ctx, err error) error {
	// pengecekan error jika error sama dengan ErrNotFound
	if errors.Is(err, ErrNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
		})
	}else if valida	, ok := err.(validator.ValidationErrors); ok {
		// jika error dari validation maka akan di tampilkan error 400
		// melewati function HandlerNotFound
		return HandlerNotFound(c, valida)

		// ini adalah custom error langsung
	}else if  customError, ok := err.(HandlerCustomErrorInterface); ok{
		return c.Status(customError.GetStatusCode()).JSON(customError.GetMessage())
	}

	// pengecekan error jika error sama dengan selain yang diatas maka akan di tampilkan error 500
	// ini sebenernya ketika yang diatas salah semua maka akan di tampilkan error 500
	// default pesannya adalah Internal Server Error
	// tapi jika app mode nya development maka akan di tampilkan error asli
	message := ErrInternalServerError.Error()
	if config.AppMode == "development"{
		message = err.Error()
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": message,
	})
}

// handler untuk menangani error validation
func HandlerNotFound(c *fiber.Ctx, err error) error {
	validationErrors := []ErrorResponse{}
	for _, err := range err.(validator.ValidationErrors) {
		// kumpulan error validation
		var elem ErrorResponse

		elem.Field = err.Field() // Export struct field name
		elem.Tag = err.Tag()           // Export struct tag
		elem.Value = err.Value()       // Export field value

		// custom message
		// pengecekannya jika tag nya ada di TagCustomMessage maka akan di tampilkan custom message
		// jika tidak maka akan di tampilkan error asli
		if function, ok := TagCustomMessage[elem.Tag]; ok {
			// dan mengirim parameter yang diperlukan untuk function tersebut
			// bisa di cek di file tag_custom_message.go
			elem.Message = function(err.Field())
		} else {
			elem.Message = err.Error()
		}

		validationErrors = append(validationErrors, elem)
	}

	// return response 400
	return c.Status(400).JSON(fiber.Map{
		"message": ErrBadParamInput.Error(),
		"data": validationErrors,
	})
}