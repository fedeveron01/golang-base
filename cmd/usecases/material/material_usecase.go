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
		return entities.Material{}, core_errors.NewInternalServerError("Material name must be at least 2 characters")
	}
	if material.Price < 0 {
		return entities.Material{}, core_errors.NewInternalServerError("Material price must be greater than or equal to 0")
	}
	if material.Stock < 0 {
		return entities.Material{}, core_errors.NewInternalServerError("Material stock must be greater than or equal to 0")
	}
	repeated := i.materialGateway.FindByName(material.Name)
	if repeated != nil {
		return entities.Material{}, core_errors.NewInternalServerError("Material already exists")
	}
	materialType := i.materialTypeGateway.FindById(material.MaterialType.ID)
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
func (i *Implementation) UpdateMaterial(material entities.Material) (entities.Material, error) {
	if material.ID <= 0 {
		return entities.Material{}, core_errors.NewInternalServerError("material id is required")
	}
	/*found := i.materialGateway.FindById(material.ID)
	if found == nil && found.ID != materialType.ID {
		return entities.MaterialType{}, core_errors.NewInternalServerError("material not exist")
	}*/
	/*repeated := i.materialTypeGateway.FindByName(materialType.Name)
	if repeated != nil && repeated.ID != materialType.ID {
		return entities.MaterialType{}, core_errors.NewInternalServerError("materialType already exists")
	}*/
	if material.Name == "" {
		return entities.Material{}, core_errors.NewBadRequestError("material name is required")
	}
	if material.Description == "" {
		return entities.Material{}, core_errors.NewBadRequestError("material description is required")
	}
	if material.Price < 0 {
		return entities.Material{}, core_errors.NewBadRequestError("material price must be greater than or equal to 0")
	}
	if material.Stock < 0 {
		return entities.Material{}, core_errors.NewBadRequestError("material stock must be greater than or equal to 0")
	}
	materialType := i.materialTypeGateway.FindById(material.MaterialType.ID)
	if materialType == nil {
		return entities.Material{}, core_errors.NewInternalServerError("Material Type not found")
	}
	if material.RepositionPoint < 0 {
		return entities.Material{}, core_errors.NewInternalServerError("Material reposition point must be greater than or equal to 0")
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
