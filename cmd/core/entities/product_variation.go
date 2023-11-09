package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type ProductVariation struct {
	core.EntitiesBase
	Number float64
	Stock  float64
}
