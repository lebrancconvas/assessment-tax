package form

type RequestPersonalDeductionData struct {
		Amount float64 `json:"amount"`
	}

	type ResponsePersonalDeductionData struct {
		PersonalDeduction float64 `json:"personalDeduction"`
	}
