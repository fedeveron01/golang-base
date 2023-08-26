package entities

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name             string
	DNI              string
	UserId           uint
	User             User
	ProductionOrders []ProductionOrder
	PurchaseOrders   []PurchaseOrder
}
