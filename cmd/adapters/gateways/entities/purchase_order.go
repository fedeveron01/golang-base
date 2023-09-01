package gateway_entities

import "gorm.io/gorm"

type PurchaseOrder struct {
	gorm.Model
	Number               int
	Description          string
	PurchaseOrderDetails []PurchaseOrderDetail
	EmployeeId           uint
}
