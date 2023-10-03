package gateway_entities

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Name            string
	Description     string
	Price           float64
	Stock           float64
	RepositionPoint float64
	MaterialTypeId  uint
	MaterialType    MaterialType
}
