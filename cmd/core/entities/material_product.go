package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type MaterialProduct struct {
	core.EntitiesBase
	Quantity float64
	Material Material
}
