package helpers

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// APIResponse is for generating template responses
func APIResponse(message string, code int, status string, data interface{}) Response {

	return Response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// FormatValidationError func which holds errors during user input validation
func FormatValidationError(err error) string {
	var errors string

	for _, e := range err.(validator.ValidationErrors) {
		errors = e.Error()
		return errors
	}

	return errors
}
