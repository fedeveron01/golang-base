package providers

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
)

type MaterialProvider interface {
	CreateMaterial(material entities.Material) error
	FindAll() ([]entities.Material, error)
	UpdateMaterial(material entities.Material) error
	DeleteMaterial(id string) error
}
