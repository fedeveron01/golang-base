package gateway_entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name            string
	Description     string
	Price           float64
	Stock           int
	MaterialProduct []MaterialProduct
}
