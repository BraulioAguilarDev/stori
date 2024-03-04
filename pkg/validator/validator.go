package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStructure(params any) map[string]string {
	validate = validator.New(validator.WithRequiredStructEnabled())
	var errors = make(map[string]string, 0)

	if err := validate.Struct(params); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("The input %s was not valid", err.Field())
		}
	}

	return errors
}
