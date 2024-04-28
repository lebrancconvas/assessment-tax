package model

import (
	"math"

	"github.com/lebrancconvas/assessment-tax/form"
)

var allowance float64 = 60000

func CalculateTax(taxData form.RequestData) (float64, error) {
	netIncome := taxData.TotalIncome - allowance

	for _, anotherAllowance := range taxData.Allowances {
		netIncome -= anotherAllowance.Amount
	}

	var taxResult float64 = 0.0

	var rates = []form.Rate{
		{LimitIncome: 150000, TaxRate: 0},
		{LimitIncome: 500000, TaxRate: 0.10},
		{LimitIncome: 1000000, TaxRate: 0.15},
		{LimitIncome: 2000000, TaxRate: 0.20},
		{LimitIncome: math.Inf(99), TaxRate: 0.35},
	}

	for _, rate := range rates {
		if netIncome < 150000 {
			break
		}
		taxResult += (netIncome * rate.TaxRate)
		netIncome -= rate.LimitIncome
	}

	taxResult -= taxData.WHT

	return taxResult, nil
}

func CalculateTaxLevel(taxData form.RequestData) ([]form.TaxLevel, error) {
	var taxLevels []form.TaxLevel

	taxLevels = []form.TaxLevel{
		{Level: "0-150,000", Tax: 0.0},
		{Level: "150,001-500,000", Tax: 0.0},
		{Level: "500,001-1,000,000", Tax: 0.0},
		{Level: "1,000,001-2,000,000", Tax: 0.0},
		{Level: "2,000,001 ขึ้นไป", Tax: 0.0},
	}

	return taxLevels, nil
}
