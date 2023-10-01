package material_type_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
)

type MaterialTypeUseCase interface {
	CreateMaterialType(materialType entities.MaterialType) error
}

type MaterialTypeTypeGateway interface {
	CreateMaterialType(materialTypeType entities.MaterialType) error
	FindByName(name string) *entities.MaterialType
}

type Implementation struct {
	materialTypeGateway MaterialTypeTypeGateway
}

func NewMaterialTypeUsecase(materialTypeTypeGateway MaterialTypeTypeGateway) *Implementation {
	return &Implementation{
		materialTypeGateway: materialTypeTypeGateway,
	}
}

func (i *Implementation) CreateMaterialType(materialType entities.MaterialType) error {
	if materialType.Name == "" {
		return core_errors.NewInternalServerError("materialType name is required")
	}
	if len(materialType.Name) < 2 {
		return core_errors.NewInternalServerError("materialType name must be at least 2 characters")
	}
	if len(materialType.Name) > 20 {
		return core_errors.NewInternalServerError("materialType name must be at most 20 characters")
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
