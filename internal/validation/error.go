package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	errorList []string
}

func (ve *ValidationError) Error() string {
	return strings.Join(ve.errorList, ",")
}

func (ve *ValidationError) Add(err string) {
	ve.errorList = append(ve.errorList, err)
}

func NewError(err error) *ValidationError {
	ValidationErrors := &ValidationError{}
	errorItem, ok := err.(validator.ValidationErrors)
	if !ok {
		ValidationErrors.Add("validation error")
		return ValidationErrors
	}

	for _, err := range errorItem {
		switch err.Tag() {
		case "required":
			ValidationErrors.Add(fmt.Sprintf("%s is required", err.Field()))
		case "email":
			ValidationErrors.Add(fmt.Sprintf("%s is not valid email", err.Field()))
		case "min":
			ValidationErrors.Add(fmt.Sprintf("%s value must be more than %s", err.Field(), err.Param()))
		}
	}
	return ValidationErrors
}
