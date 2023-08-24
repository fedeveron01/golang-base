package entities

import (
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	TotalInstallment float64
	InstallmentSize  float64
	InstallmentCount int
	ExpirationDay    int
}
