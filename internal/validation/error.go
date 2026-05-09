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

func NewError(err error) error {
	ValidationErrors := &ValidationError{}
	errorList, ok := err.(validator.ValidationErrors) // comma ok idiom
	if !ok {
		ValidationErrors.Add("validation error")
		return ValidationErrors
	}

	for _, err := range errorList {
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

func IsErrValidation(err error) bool {
	if _, ok := err.(*ValidationError); !ok {
		return false
	}
	return true
}
