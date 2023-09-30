package gateway_entities

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Name            string
	Description     string
	Price           float64
	Stock           int
	MaterialTypeId  uint
	MaterialType    MaterialType
	RepositionPoint int
}
