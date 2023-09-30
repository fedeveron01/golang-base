package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type MaterialType struct {
	core.EntitiesBase
	Name        string
	Description string
}
