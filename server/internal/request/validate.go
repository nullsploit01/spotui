package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Errors map[string]string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Validation failed for %d fields", len(v.Errors))
}

func ValidateRequest(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, fieldError := range validationErrors {
			errors[fieldError.Field()] = getErrorMessage(fieldError)
		}
		return &ValidationError{Errors: errors}
	}

	return err
}

func getErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Value is too short, minimum is %s", fieldError.Param())
	case "max":
		return fmt.Sprintf("Value is too long, maximum is %s", fieldError.Param())
	case "gte":
		return fmt.Sprintf("Value is too small, minimum is %s", fieldError.Param())
	case "lte":
		return fmt.Sprintf("Value is too large, maximum is %s", fieldError.Param())
	default:
		return "Invalid value"
	}
}
