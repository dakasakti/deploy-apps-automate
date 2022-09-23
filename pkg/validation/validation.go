package validation

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, MessageValidate(err))
	}
	return nil
}

func MessageValidate(err error) map[string]string {
	messages := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			messages[err.Field()] = fmt.Sprintf("%s is required", err.Field())
		case "lowercase":
			messages[err.Field()] = fmt.Sprintf("%s must be lowercase", err.Field())
		case "min":
			messages[err.Field()] = fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())
		case "max":
			messages[err.Field()] = fmt.Sprintf("%s must be less than %s characters", err.Field(), err.Param())
		}
	}

	return messages
}
