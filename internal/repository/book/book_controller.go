package book

import (
	"strconv"

	"github.com/dakasakti/deploy-apps-hexagonal/internal/middlewares"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/models"
	"github.com/labstack/echo/v4"
)

var (
	datas []models.Book
)

func GetBooks(c echo.Context) error {
	if len(datas) == 0 {
		return c.JSON(404, echo.Map{
			"message": "data tidak ditemukan",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil get all data book",
		"data":    datas,
	})
}

func GetBook(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "id yang dimasukkan salah",
		})
	}

	for _, val := range datas {
		if val.ID == uint(conv_id) {
			book := models.Book{
				ID:       val.ID,
				Title:    val.Title,
				Penulis:  val.Penulis,
				Penerbit: val.Penerbit,
			}

			return c.JSON(200, echo.Map{
				"message": "berhasil get data book",
				"data":    book,
			})
		}
	}

	return c.JSON(404, echo.Map{
		"message": "data tidak ditemukan",
	})
}

func CreateBook(c echo.Context) error {
	user_id := middlewares.ExtractToken(c)
	var req models.BookRequest
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

	book := models.Book{
		ID:       uint(len(datas) + 1),
		Title:    req.Title,
		Penulis:  req.Penulis,
		Penerbit: req.Penerbit,
		UserID:   user_id,
	}

	datas = append(datas, book)

	return c.JSON(200, echo.Map{
		"message": "berhasil membuat data book",
		"data":    book,
	})
}

func UpdateBook(c echo.Context) error {
	user_id := middlewares.ExtractToken(c)
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "id yang dimasukkan salah",
		})
	}

	var req models.BookUpdateRequest
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "data yang dimasukkan salah",
		})
	}

	if req.Title == "" && req.Penulis == "" && req.Penerbit == "" {
		return c.JSON(204, echo.Map{
			"message": "tidak ada yang diperbaiki",
		})
	}

	err = c.Validate(req)
	if err != nil {
		return c.JSON(400, err)
	}

	book := models.Book{
		ID:       uint(conv_id),
		Title:    req.Title,
		Penulis:  req.Penulis,
		Penerbit: req.Penerbit,
		UserID:   user_id,
	}

	for _, val := range datas {
		if val.ID == uint(conv_id) && val.UserID == user_id {
			datas[conv_id-1] = book
		} else {
			return c.JSON(403, echo.Map{
				"message": "akses tidak diperbolehkan",
			})
		}
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil update data book",
		"data":    book,
	})
}

func DeleteBook(c echo.Context) error {
	user_id := middlewares.ExtractToken(c)
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "id yang dimasukkan salah",
		})
	}

	for _, val := range datas {
		if val.ID == uint(conv_id) && val.UserID == user_id {
			datas = append(datas[:conv_id-1], datas[conv_id:]...)
		} else {
			return c.JSON(403, echo.Map{
				"message": "akses tidak diperbolehkan",
			})
		}
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil menghapus data book",
	})
}
