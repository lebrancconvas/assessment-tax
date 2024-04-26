package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lebrancconvas/assessment-tax/form"
	"github.com/lebrancconvas/assessment-tax/model"
)

func CalculateTaxController(c echo.Context) error {
	var requestData form.RequestData

	err := c.Bind(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := model.CalculateTax(requestData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, form.ResponseData{
		Tax: result,
	})
}
