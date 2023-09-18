package gateway_entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	Password string
	Inactive bool `gorm:"default:false"`
}
