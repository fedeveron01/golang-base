package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type ProductionOrderDetail struct {
	core.EntitiesBase
	Quantity int
	Product  Product
}
