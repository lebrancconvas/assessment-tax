package server

import (
	"github.com/labstack/echo/v4"
	"github.com/lebrancconvas/assessment-tax/controller"
)

func NewRouter() *echo.Echo {
	router := echo.New()

	// Router
	router.GET("/", controller.Hello)

	// Core API
	router.POST("/api/tax", controller.CalculateTaxController)

	return router
}
