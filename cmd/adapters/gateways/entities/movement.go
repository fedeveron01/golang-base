package gateway_entities

import "gorm.io/gorm"

type Movement struct {
	gorm.Model
	Quantity           float64
	MaterialId         uint
	Material           Material
	MovementTypeId     uint
	ProductVariation   ProductVariation
	ProductVariationId uint
}
