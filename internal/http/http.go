package http

import (
	"github.com/dakasakti/deploy-apps-hexagonal/internal/app/user"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/factory"
	"github.com/dakasakti/deploy-apps-hexagonal/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	e.GET("api/v1/status", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"status": "OK",
		})
	})

	v1 := e.Group("/api/v1")
	user.NewUserController(f).Route(v1)
}
