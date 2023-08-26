package entities

import "gorm.io/gorm"

type MeasurementUnit struct {
	gorm.Model
	Name string
}
