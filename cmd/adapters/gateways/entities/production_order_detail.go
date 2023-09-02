package gateway_entities

import "gorm.io/gorm"

type ProductionOrderDetail struct {
	gorm.Model
	Quantity          int
	ProductID         uint
	Product           Product
	ProductionOrderID uint
}
