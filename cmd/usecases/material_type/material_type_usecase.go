package material_type_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
)

type MaterialTypeUseCase interface {
	FindAll() ([]entities.MaterialType, error)
	CreateMaterialType(materialType entities.MaterialType) error
	UpdateMaterialType(materialType entities.MaterialType) (entities.MaterialType, error)
	DeleteMaterialType(id uint) error
}

type MaterialTypeTypeGateway interface {
	FindAll() ([]entities.MaterialType, error)
	FindById(id uint) *entities.MaterialType
	FindByName(name string) *entities.MaterialType
	CreateMaterialType(materialTypeType entities.MaterialType) error
	UpdateMaterialType(materialTypeType entities.MaterialType) (entities.MaterialType, error)
	DeleteMaterialType(id uint) error
}

type Implementation struct {
	materialTypeGateway MaterialTypeTypeGateway
}

func NewMaterialTypeUsecase(materialTypeTypeGateway MaterialTypeTypeGateway) *Implementation {
	return &Implementation{
		materialTypeGateway: materialTypeTypeGateway,
	}
}

func (i *Implementation) FindAll() ([]entities.MaterialType, error) {
	materials, err := i.materialTypeGateway.FindAll()
	if err != nil {
		return nil, err
	}
	return materials, nil
}

func (i *Implementation) CreateMaterialType(materialType entities.MaterialType) error {
	if materialType.Name == "" {
		return core_errors.NewInternalServerError("materialType name is required")
	}
	if len(materialType.Name) < 2 {
		return core_errors.NewInternalServerError("materialType name must be at least 2 characters")
	}
	repeated := i.materialTypeGateway.FindByName(materialType.Name)
	if repeated != nil {
		return core_errors.NewInternalServerError("materialType already exists")
	}
	err := i.materialTypeGateway.CreateMaterialType(materialType)
	if err != nil {
		return err
	}
	return nil
}

func (i *Implementation) UpdateMaterialType(materialType entities.MaterialType) (entities.MaterialType, error) {
	if materialType.ID <= 0 {
		return entities.MaterialType{}, core_errors.NewInternalServerError("materialType id is required")
	}

	found := i.materialTypeGateway.FindById(materialType.ID)
	if found == nil && found.ID != materialType.ID {
		return entities.MaterialType{}, core_errors.NewInternalServerError("materialType not exist")
	}

	repeated := i.materialTypeGateway.FindByName(materialType.Name)
	if repeated != nil && repeated.ID != materialType.ID {
		return entities.MaterialType{}, core_errors.NewInternalServerError("materialType already exists")
	}

	if materialType.Name == "" {
		return entities.MaterialType{}, core_errors.NewBadRequestError("materialType name is required")
	}

	if len(materialType.Name) < 2 {
		return entities.MaterialType{}, core_errors.NewBadRequestError("materialType name must be at least 2 characters")
	}
	if len(materialType.Name) > 20 {
		return entities.MaterialType{}, core_errors.NewBadRequestError("materialType name must be less than 20 characters")
	}

	materialType, err := i.materialTypeGateway.UpdateMaterialType(materialType)
	if err != nil {
		return entities.MaterialType{}, err
	}
	return materialType, nil

}

func (i *Implementation) DeleteMaterialType(id uint) error {
	err := i.materialTypeGateway.DeleteMaterialType(id)
	if err != nil {
		return err
	}
	return nil
}
