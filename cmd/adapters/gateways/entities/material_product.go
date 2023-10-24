package gateway_entities

import "gorm.io/gorm"

type MaterialProduct struct {
	gorm.Model
	Quantity   int
	MaterialId uint
	Material   Material
	ProductId  uint
	Product    Product
}
