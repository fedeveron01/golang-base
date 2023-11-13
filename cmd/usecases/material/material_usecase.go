package material_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
)

type MaterialUseCase interface {
	CreateMaterial(material entities.Material) (entities.Material, error)
	FindAll() ([]entities.Material, error)
	UpdateMaterial(material entities.Material) (entities.Material, error)
	DeleteMaterial(id string) error
}

type MaterialTypeGateway interface {
	FindById(id uint) *entities.MaterialType
}
type MaterialGateway interface {
	FindAll() ([]entities.Material, error)
	FindById(id uint) *entities.Material
	FindByName(name string) *entities.Material
	CreateMaterial(material entities.Material) (entities.Material, error)
	UpdateMaterial(material entities.Material) (entities.Material, error)
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
	if len(material.Name) < 2 {
		return entities.Material{}, core_errors.NewInternalServerError("material name must be at least 2 characters")
	}
	if len(material.Name) > 30 {
		return entities.Material{}, core_errors.NewInternalServerError("material name must be greater than 30 characters")
	}
	if material.Price < 0 {
		return entities.Material{}, core_errors.NewInternalServerError("material price must be greater than or equal to 0")
	}
	if material.Stock < 0 {
		return entities.Material{}, core_errors.NewInternalServerError("material stock must be greater than or equal to 0")
	}
	repeated := i.materialGateway.FindByName(material.Name)
	if repeated != nil {
		return entities.Material{}, core_errors.NewInternalServerError("material already exists")
	}
	materialType := i.materialTypeGateway.FindById(material.MaterialType.ID)
	if materialType == nil {
		return entities.Material{}, core_errors.NewInternalServerError("material Type not found")
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
func (i *Implementation) UpdateMaterial(material entities.Material) (entities.Material, error) {
	if material.ID <= 0 {
		return entities.Material{}, core_errors.NewBadRequestError("material id is required")
	}
	found := i.materialGateway.FindById(material.ID)
	if found == nil || found.ID != material.ID {
		return entities.Material{}, core_errors.NewInternalServerError("material not exist")
	}

	if material.Name == "" {
		return entities.Material{}, core_errors.NewBadRequestError("material name is required")
	}
	if len(material.Name) < 2 {
		return entities.Material{}, core_errors.NewInternalServerError("material name must be at least 2 characters")
	}
	if len(material.Name) > 30 {
		return entities.Material{}, core_errors.NewInternalServerError("material name must be greater than 30 characters")
	}
	if material.Price < 0 {
		return entities.Material{}, core_errors.NewBadRequestError("material price must be greater than or equal to 0")
	}
	if material.Stock < 0 {
		return entities.Material{}, core_errors.NewBadRequestError("material stock must be greater than or equal to 0")
	}
	materialType := i.materialTypeGateway.FindById(material.MaterialType.ID)
	if materialType == nil {
		return entities.Material{}, core_errors.NewInternalServerError("material Type not found")
	}
	if material.RepositionPoint < 0 {
		return entities.Material{}, core_errors.NewInternalServerError("material reposition point must be greater than or equal to 0")
	}
	material, err := i.materialGateway.UpdateMaterial(material)
	if err != nil {
		return entities.Material{}, err
	}
	return material, nil
}
func (i *Implementation) DeleteMaterial(id string) error {
	return i.materialGateway.DeleteMaterial(id)
}
