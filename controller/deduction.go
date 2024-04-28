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

	if req.Amount < 60000 || req.Amount > 100000 {
		return c.JSON(http.StatusBadRequest, form.Message{
			Message: "Personal deduction must be between 60,000 and 100,000.",
		})
	}

	res, err := model.GetPersonalDeduction()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, form.ResponsePersonalDeductionData{
		PersonalDeduction: res,
	})
}
