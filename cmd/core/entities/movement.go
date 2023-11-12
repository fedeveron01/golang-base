package entities

import (
	"github.com/fedeveron01/golang-base/cmd/core"
	"time"
)

type Movement struct {
	core.EntitiesBase
	Number             int
	Type               string
	Total              float64
	DateTime           time.Time
	Description        string
	IsMaterialMovement bool
	MovementDetail     []MovementDetail
}
