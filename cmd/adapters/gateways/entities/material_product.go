package gateway_entities

import "gorm.io/gorm"

type MaterialProduct struct {
	gorm.Model
	Quantity   float64
	MaterialId uint
	Material   Material
	ProductId  uint
	Product    Product
}
