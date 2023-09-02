package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type Employee struct {
	core.EntitiesBase
	Name             string
	LastName         string
	DNI              string
	User             User
	Charge           Charge
	ProductionOrders []ProductionOrder
	PurchaseOrders   []PurchaseOrder
}
