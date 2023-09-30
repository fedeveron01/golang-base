package material_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
)

type MaterialUseCase interface {
	CreateMaterial(material entities.Material) (entities.Material, error)
	FindAll() ([]entities.Material, error)
	UpdateMaterial(material entities.Material) error
	DeleteMaterial(id string) error
}

type MaterialTypeGateway interface {
	FindByName(name string) *entities.MaterialType
}
type MaterialGateway interface {
	FindAll() ([]entities.Material, error)
	FindByName(name string) *entities.Material
	CreateMaterial(material entities.Material) (entities.Material, error)
	UpdateMaterial(material entities.Material) error
	DeleteMaterial(id string) error
}

type Implementation struct {
	materialGateway     MaterialGateway
	materialTypeGateway MaterialTypeGateway
}

func NewMaterialUsecase(materialGateway MaterialGateway, materialTypeGateway MaterialTypeGateway) *Implementation {
	return &Implementation{
		materialGateway:     materialGateway,
		materialTypeGateway: materialTypeGateway,
	}
}

func (i *Implementation) CreateMaterial(material entities.Material) (entities.Material, error) {
	repeated := i.materialGateway.FindByName(material.Name)
	if repeated != nil {
		return entities.Material{}, core_errors.NewInternalServerError("Material already exists")
	}
	materialType := i.materialTypeGateway.FindByName(material.MaterialType.Name)
	if materialType == nil {
		return entities.Material{}, core_errors.NewInternalServerError("Material Type not found")
	}

	material, err := i.materialGateway.CreateMaterial(material)
	if err != nil {
		return entities.Material{}, err
	}
	return material, err
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
