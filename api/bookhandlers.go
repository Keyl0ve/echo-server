package api

import (
	"net/http"

	"github.com/Keyl0ve/echo-server/model"

	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	book := new(model.Book)
	if err := c.Bind(book); err != nil {
		return err
	}

	err := model.SaveBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, book)
}

func GetBooksList(c echo.Context) error {
	books, err := model.GetBooksList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}
