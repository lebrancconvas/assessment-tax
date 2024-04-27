package form

type RequestData struct {
	TotalIncome float64 `json:"totalIncome"`
	WHT	float64 `json:"wht"`
	Allowances []Allowance `json:"allowances"`
}

type Allowance struct {
	AllowanceType string `json:"allowanceType"`
	Amount float64 `json:"amount"`
}

type ResponseData struct {
	Tax float64 `json:"tax"`
	TaxLevels []TaxLevel `json:"taxLevel"`
}

type TaxLevel struct {
	Level string `json:"level"`
	Tax float64 `json:"tax"`
}

type Rate struct {
	LimitIncome float64
	TaxRate float64
}
