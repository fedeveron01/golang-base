package entities

import "gorm.io/gorm"

type MaterialType struct {
	gorm.Model
	Name        string
	Description string
}
