package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

func GetValidationErrorMessage(err validator.FieldError) string {
	fieldName := strings.ToLower(err.Field())

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s field is required.", fieldName)
	case "min":
		return fmt.Sprintf("%s should be at least %s characters long.", fieldName, err.Param())
	case "max":
		return fmt.Sprintf("%s should be at most %s characters long.", fieldName, err.Param())
	case "email":
		return "Enter a valid email address."
	case "gte":
		return fmt.Sprintf("%s should be greater than %s.", fieldName, err.Param())
	case "oneof":
		return fmt.Sprintf("%s should be one of the allowed values: %s.", fieldName, err.Param())
		// Add more cases for other validation tags as needed.

	default:
		return fmt.Sprintf("Invalid input: %s.", err.Param())
	}
}
