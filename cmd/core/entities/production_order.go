package entities

import (
	"gorm.io/gorm"
	"time"
)

type ProductionOrder struct {
	gorm.Model
	StartDateTime         time.Time
	EndDateTime           time.Time
	ProductionOrderDetail []ProductionOrderDetail
	EmployeeID            uint
}
