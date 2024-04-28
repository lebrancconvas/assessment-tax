package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lebrancconvas/assessment-tax/form"
	"github.com/lebrancconvas/assessment-tax/model"
)

func SetPersonalDeduction(c echo.Context) error {
	req := form.RequestPersonalDeductionData{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = model.CreatePersonalDeduction(req.Amount)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	res, err := model.GetPersonalDeduction()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, form.ResponsePersonalDeductionData{
		PersonalDeduction: res,
	})
}
