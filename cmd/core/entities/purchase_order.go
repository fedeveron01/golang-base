package entities

import "gorm.io/gorm"

type PurchaseOrder struct {
	gorm.Model
	Number      int
	Description string
	EmployeeId  uint
}
