package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}

func APIResponse(message string, err []string, status bool, data interface{}) Response {

	response := Response{
		Status:  status,
		Message: message,
		Errors:  err,
		Data:    data,
	}

	return response
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
