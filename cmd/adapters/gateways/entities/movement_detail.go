package gateway_entities

import "gorm.io/gorm"

type MovementDetail struct {
	gorm.Model
	Quantity           float64
	Price              float64
	Material           *Material
	MaterialId         uint
	ProductVariation   *ProductVariation
	ProductVariationId uint
	MovementId         uint
}
