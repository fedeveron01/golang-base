package entities

import (
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
)

type MaterialType struct {
	core.EntitiesBase
	Name              string
	Description       string
	UnitOfMeasurement enums.Enum
}
