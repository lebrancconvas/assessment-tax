package model

import (
	"testing"

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
