package gateway_entities

import "gorm.io/gorm"

type ProductVariation struct {
	gorm.Model
	Number    float64
	Stock     float64
	ProductID uint
	Product   Product
}
