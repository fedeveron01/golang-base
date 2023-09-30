package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type Material struct {
	core.EntitiesBase
	Name            string
	Description     string
	Price           float64
	Stock           int
	RepositionPoint int
	MaterialType    MaterialType
	MeasurementUnit MeasurementUnit
}
