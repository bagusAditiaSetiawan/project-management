package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error       bool        `json:"error"`
	FailedField string      `json:"failedField"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
	Message     string      `json:"message"`
}

func CustomErrorValidator(errs validator.ValidationErrors) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	if errs != nil {
		for _, err := range errs {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true
			elem.Message = fmt.Sprintf("%s is %s", elem.FailedField, elem.Tag)

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
