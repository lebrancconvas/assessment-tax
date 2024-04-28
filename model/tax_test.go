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

	t.Run("Test Calculate Tax from Total Income with WHT", func(t *testing.T) {
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

		assert.EqualValues(t, 4000.0, taxResult)

	})

	t.Run("Test Calculate Tax from Total Income with WHT and allowances", func(t *testing.T) {
		request := form.RequestData{
			TotalIncome: 50000.0,
			WHT: 0.0,
			Allowances: []form.Allowance{
				{
					AllowanceType: "donation",
					Amount: 200000.0,
				},
			},
		}

		taxResult, _ := CalculateTax(request)

		assert.EqualValues(t, 19000.0, taxResult)

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

func TestCalculateTaxWithWHTServer(t *testing.T) {
	testRequestJSON := `{"totalIncome: 50000.0, "wht": 25000.0, "allowances": [{"allowanceType": "donation", "amount": 0.0}]`
	testResponseJSON := `{tax: 4000.0}`

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

func TestCalculateTaxWithWHTandAllowancesServer(t *testing.T) {
	testRequestJSON := `{"totalIncome: 50000.0, "wht": 0.0, "allowances": [{"allowanceType": "donation", "amount": 200000.0}]`
	testResponseJSON := `{tax: 19000.0}`

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

func TestCalculateTaxWithStepResponse(t *testing.T) {
	testRequestJSON := `{"totalIncome: 50000.0, "wht": 0.0, "allowances": [{"allowanceType": "donation", "amount": 200000.0}]`
	testResponseJSON := `
		{
			"tax": 19000.0,
			"taxLevel": [
				{
					"level": "0-150,000",
					"tax": 0.0
				},
				{
					"level": "150,001-500,000",
					"tax": 19000.0
				},
				{
					"level": "500,001-1,000,000",
					"tax": 0.0
				},
				{
					"level": "1,000,001-2,000,000",
					"tax": 0.0
				},
				{
					"level": "2,000,001 ขึ้นไป",
					"tax": 0.0
				}
			]
		}
	`

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

func TestCalculateTaxWithAllowancesWithStepResponse(t *testing.T) {
	testRequestJSON := `{"totalIncome: 50000.0, "wht": 0.0, "allowances": [{"allowanceType": "donation", "amount": 200000.0}, {"allowanceType": "k-receipt", "amount": 100000.0}]`
	testResponseJSON := `
		{
			"tax": 19000.0,
			"taxLevel": [
				{
					"level": "0-150,000",
					"tax": 0.0
				},
				{
					"level": "150,001-500,000",
					"tax": 14000.0
				},
				{
					"level": "500,001-1,000,000",
					"tax": 0.0
				},
				{
					"level": "1,000,001-2,000,000",
					"tax": 0.0
				},
				{
					"level": "2,000,001 ขึ้นไป",
					"tax": 0.0
				}
			]
		}
	`

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
