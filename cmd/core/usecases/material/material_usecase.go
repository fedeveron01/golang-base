package material_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/providers"
)

type MaterialUsecase interface {
	CreateMaterial(material entities.Material) error
	FindAll() ([]entities.Material, error)
	UpdateMaterial(material entities.Material) error
	DeleteMaterial(id string) error
}
type Implementation struct {
	materialProvider providers.MaterialProvider
}

func NewMaterialUsecase(materialProvider providers.MaterialProvider) *Implementation {
	return &Implementation{
		materialProvider: materialProvider,
	}
}

func (i *Implementation) CreateMaterial(material entities.Material) error {
	err := i.materialProvider.CreateMaterial(material)
	if err != nil {
		return err
	}
	return nil
}
func (i *Implementation) FindAll() ([]entities.Material, error) {
	return i.materialProvider.FindAll()
}
func (i *Implementation) UpdateMaterial(material entities.Material) error {
	return i.materialProvider.UpdateMaterial(material)
}
func (i *Implementation) DeleteMaterial(id string) error {
	return i.materialProvider.DeleteMaterial(id)
}
