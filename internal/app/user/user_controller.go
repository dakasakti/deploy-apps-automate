package user

import (
	"github.com/dakasakti/deploy-apps-hexagonal/internal/factory"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/middlewares"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/models"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	Login(c echo.Context) error
}

type userController struct {
	service UserService
}

func NewUserController(f *factory.Factory) *userController {
	return &userController{
		service: NewUserService(f),
	}
}

func (uc *userController) GetUsers(c echo.Context) error {
	results, err := uc.service.GetUsers(c.Request().Context())
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil get all data user",
		"data":    results,
	})
}

func (uc *userController) GetUser(c echo.Context) error {
	user_id := middlewares.ExtractToken(c)

	result, err := uc.service.GetUser(c.Request().Context(), user_id)
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil get data user",
		"data":    result,
	})
}

func (uc *userController) CreateUser(c echo.Context) error {
	var req models.UserRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "data yang dimasukkan salah",
		})
	}

	err = c.Validate(req)
	if err != nil {
		return c.JSON(400, err)
	}

	result, err := uc.service.CreateUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(201, echo.Map{
		"message": "berhasil membuat data user",
		"data": echo.Map{
			"inserted_id": result,
		},
	})
}

func (uc *userController) UpdateUser(c echo.Context) error {
	user_id := middlewares.ExtractToken(c)
	var req models.UserUpdateRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": err.Error(),
		})
	}

	err = c.Validate(req)
	if err != nil {
		return c.JSON(400, err)
	}

	err = uc.service.UpdateUser(c.Request().Context(), user_id, req)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil update data user",
	})
}

func (uc *userController) DeleteUser(c echo.Context) error {
	user_id := middlewares.ExtractToken(c)

	err := uc.service.DeleteUser(c.Request().Context(), user_id)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil menghapus data user",
	})
}

func (uc *userController) Login(c echo.Context) error {
	var req models.UserLogin
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "data yang dimasukkan salah",
		})
	}

	err = c.Validate(req)
	if err != nil {
		return c.JSON(400, err)
	}

	result, err := uc.service.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil login",
		"token":   result,
	})
}
