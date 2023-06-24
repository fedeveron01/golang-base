package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscription struct {
	ID               primitive.ObjectID `json:"_id,omitempty"`
	TotalInstallment float64            `json:"total,omitempty"`
	Members          []Member           `json:"members,omitempty"`
	InstallmentSize  float64            `json:"installment_size,omitempty"`
	InstallmentCount int                `json:"installment_count,omitempty"`
	ExpirationDay    int                `json:"expiration_day,omitempty"`
}
