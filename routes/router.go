package routes

import (
	api "github.com/Keyl0ve/echo-server/api"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	api.SetMiddlewares(e)

	// Routes for book management
	e.POST("/books", api.CreateBook)
	e.GET("/books", api.GetBooksList)
}
