package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	router := echo.New()

	// Router
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	return router
}
