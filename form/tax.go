package form

type TaxRequest struct {
	TotalIncome float64 `json:"totalIncome"`
	WHT	float64 `json:"wht"`
	Allowances []Allowance `json:"allowances"`
}

type Allowance struct {
	AllowanceType string `json:"allowanceType"`
	Amount float64 `json:"amount"`
}
