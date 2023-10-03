package material_type_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
)

type MaterialTypeUseCase interface {
	FindAll() ([]entities.MaterialType, error)
	CreateMaterialType(materialType entities.MaterialType) error

	DeleteMaterialType(id uint) error
}

type MaterialTypeTypeGateway interface {
	FindAll() ([]entities.MaterialType, error)
	FindByName(name string) *entities.MaterialType
	CreateMaterialType(materialTypeType entities.MaterialType) error
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

func (i *Implementation) DeleteMaterialType(id uint) error {
	err := i.materialTypeGateway.DeleteMaterialType(id)
	if err != nil {
		return err
	}
	return nil
}
