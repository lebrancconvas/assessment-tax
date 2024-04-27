package model

import (
	"math"

	"github.com/lebrancconvas/assessment-tax/form"
)

var allowance float64 = 60000

func CalculateTax(taxData form.RequestData) (float64, error) {
	netIncome := taxData.TotalIncome - allowance

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
