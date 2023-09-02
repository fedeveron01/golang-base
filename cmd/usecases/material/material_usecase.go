package material_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

type MaterialUsecase interface {
	CreateMaterial(material entities.Material) error
	FindAll() ([]entities.Material, error)
	UpdateMaterial(material entities.Material) error
	DeleteMaterial(id string) error
}

type MaterialGateway interface {
	CreateMaterial(material entities.Material) error
	FindAll() ([]entities.Material, error)
	UpdateMaterial(material entities.Material) error
	DeleteMaterial(id string) error
}

type Implementation struct {
	materialGateway MaterialGateway
}

func NewMaterialUsecase(materialGateway MaterialGateway) *Implementation {
	return &Implementation{
		materialGateway: materialGateway,
	}
}

func (i *Implementation) CreateMaterial(material entities.Material) error {
	err := i.materialGateway.CreateMaterial(material)
	if err != nil {
		return err
	}
	return nil
}
func (i *Implementation) FindAll() ([]entities.Material, error) {
	return i.materialGateway.FindAll()
}
func (i *Implementation) UpdateMaterial(material entities.Material) error {
	return i.materialGateway.UpdateMaterial(material)
}
func (i *Implementation) DeleteMaterial(id string) error {
	return i.materialGateway.DeleteMaterial(id)
}
