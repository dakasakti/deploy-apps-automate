package user

import (
	"github.com/dakasakti/deploy-apps-hexagonal/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func (c *userController) Route(e *echo.Group) {
	// users
	users := e.Group("/users")
	users.GET("", c.GetUsers, middlewares.JWTSign())
	users.POST("", c.CreateUser)

	//profile
	profile := e.Group("/users/profile", middlewares.JWTSign())
	profile.GET("", c.GetUser)
	profile.PUT("", c.UpdateUser)
	profile.DELETE("", c.DeleteUser)

	// login
	e.POST("/login", c.Login)
}
