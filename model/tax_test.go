package model

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lebrancconvas/assessment-tax/controller"
	"github.com/lebrancconvas/assessment-tax/form"
	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	t.Run("Test Calculate Tax from Total Income", func(t *testing.T) {
		request := form.RequestData{
			TotalIncome: 50000.0,
			WHT: 0.0,
			Allowances: []form.Allowance{
				{
					AllowanceType: "donation",
					Amount: 0.0,
				},
			},
		}

		taxResult, _ := CalculateTax(request)

		assert.EqualValues(t, 29000.0, taxResult)

	})
}

func TestCalculateTaxServer(t *testing.T) {
	testRequestJSON := `{"totalIncome: 50000.0, "wht": 0.0, "allowances": [{"allowanceType": "donation", "amount": 0.0}]`
	testResponseJSON := `{tax: 29000.0}`

	router := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/tax/calculations", strings.NewReader(testRequestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := router.NewContext(req, rec)

	if assert.NoError(t, controller.CalculateTaxController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, testResponseJSON, rec.Body.String())
	}
}
