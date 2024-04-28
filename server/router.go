package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lebrancconvas/assessment-tax/controller"
	"github.com/lebrancconvas/assessment-tax/utils"
)

func NewRouter() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	api := router.Group("")
	api.Use(middleware.BasicAuth(utils.Authenticated))

	// Router
	api.GET("/", controller.Hello)

	// Core API
	api.POST("/tax/calculations", controller.CalculateTaxController)

	// router.POST("/admin/deductions/personal")

	return router
}
