package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type MovementDetail struct {
	core.EntitiesBase
	Quantity         float64
	Price            float64
	Material         *Material
	ProductVariation *ProductVariation
}
