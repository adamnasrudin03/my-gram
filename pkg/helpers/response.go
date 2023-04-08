package helpers

import (
	"github.com/go-playground/validator/v10"
)

type ResponseDefault struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// APIResponse is for generating template responses
func APIResponse(message string, code int, err string) ResponseDefault {

	return ResponseDefault{
		Code:    code,
		Message: message,
		Error:   err,
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
