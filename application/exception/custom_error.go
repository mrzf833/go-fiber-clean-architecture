package exception

import "github.com/gofiber/fiber/v2"

type HandlerCustomError struct {
	Message fiber.Map
	StatusCode int
}

type HandlerCustomErrorInterface interface {
	Error() string
	GetMessage() fiber.Map
	GetStatusCode() int
}

func NewHandlerCustomError(statusError int, jsonError fiber.Map) HandlerCustomErrorInterface {
	return &HandlerCustomError{
		Message: jsonError,
		StatusCode: statusError,
	}
}

func (st HandlerCustomError)GetMessage() fiber.Map {
	return st.Message
}

func (st HandlerCustomError)GetStatusCode() int {
	return st.StatusCode
}

func (st HandlerCustomError)Error() string {
	return "Custom Error"
}
