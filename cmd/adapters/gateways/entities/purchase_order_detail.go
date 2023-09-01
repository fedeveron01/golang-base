package gateway_entities

import "gorm.io/gorm"

type PurchaseOrderDetail struct {
	gorm.Model
	Quantity        int
	MaterialID      uint
	Material        Material
	PurchaseOrderID uint
}
