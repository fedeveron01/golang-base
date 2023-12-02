package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type Product struct {
	core.EntitiesBase
	Name             string
	Description      string
	Color            string
	Size             float64
	ImageUrl         string
	Price            float64
	Stock            int
	MaterialProduct  []MaterialProduct
	ProductVariation []ProductVariation
}
