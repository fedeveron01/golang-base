package gateway_entities

import "gorm.io/gorm"

type Charge struct {
	gorm.Model
	Name string
}
