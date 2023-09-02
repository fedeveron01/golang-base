package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type PurchaseOrderDetail struct {
	core.EntitiesBase
	Quantity int
	Material Material
}
