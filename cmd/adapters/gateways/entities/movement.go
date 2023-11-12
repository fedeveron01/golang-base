package gateway_entities

import (
	"gorm.io/gorm"
	"time"
)

type Movement struct {
	gorm.Model
	Number             float64
	Type               string
	Total              float64
	DateTime           time.Time
	Description        string
	MovementDetail     []MovementDetail
	IsMaterialMovement bool
	EmployeeId         uint
}
