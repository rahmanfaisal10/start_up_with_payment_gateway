package response

import "github.com/go-playground/validator/v10"

func ErrorValidationResponse(err error) []string {
	errors := make([]string, 0)
	for _, v := range err.(validator.ValidationErrors) {
		errors = append(errors, v.Error())
	}
	return errors
}
