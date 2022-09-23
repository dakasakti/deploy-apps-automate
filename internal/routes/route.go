package routes

import (
	"net/http"

	"github.com/dakasakti/deploy-apps-hexagonal/internal/app/book"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// books
	books := e.Group("/books")
	books.GET("", book.GetBooks)
	books.POST("", book.CreateBook, middlewares.JWTSign())

	book_id := books.Group("/:id")
	book_id.GET("", book.GetBook)
	book_id.PUT("", book.UpdateBook, middlewares.JWTSign())
	book_id.DELETE("", book.DeleteBook, middlewares.JWTSign())

	return e
}
