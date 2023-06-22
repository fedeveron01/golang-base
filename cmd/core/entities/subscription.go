package entities

type Subscription struct {
	TotalInstallment float64  `json:"total"`
	Members          []Member `json:"members"`
	InstallmentSize  float64  `json:"installment_size"`
	InstallmentCount int      `json:"installment_count"`
	ExpirationDay    int      `json:"expirationDay"`
}
