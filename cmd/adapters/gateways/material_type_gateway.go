package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type MaterialTypeGatewayImpl struct {
	materialTypeRepository repositories.MaterialTypeRepository
}

func NewMaterialTypeGateway(materialTypeRepository repositories.MaterialTypeRepository) *MaterialTypeGatewayImpl {
	return &MaterialTypeGatewayImpl{
		materialTypeRepository: materialTypeRepository,
	}
}

func (e *MaterialTypeGatewayImpl) FindAll() ([]entities.MaterialType, error) {
	materialTypesDB, err := e.materialTypeRepository.FindAll()
	if err != nil {
		return nil, err
	}
	materialTypes := make([]entities.MaterialType, len(materialTypesDB))
	for i, materialTypeDB := range materialTypesDB {
		materialTypes[i] = e.ToBusinessEntity(materialTypeDB)
	}
	return materialTypes, err
}

func (e *MaterialTypeGatewayImpl) FindByName(name string) *entities.MaterialType {
	materialTypeDB := e.materialTypeRepository.FindByName(name)
	if materialTypeDB == nil {
		return nil
	}
	materialType := e.ToBusinessEntity(*materialTypeDB)
	return &materialType
}

func (e *MaterialTypeGatewayImpl) CreateMaterialType(materialType entities.MaterialType) error {
	materialTypeDB := e.ToServiceEntity(materialType)

	return e.materialTypeRepository.CreateMaterialType(materialTypeDB)
}

func (e *MaterialTypeGatewayImpl) UpdateMaterialType(materialType entities.MaterialType) error {

	materialTypeDB := e.ToServiceEntity(materialType)
	return e.materialTypeRepository.UpdateMaterialType(materialTypeDB)
}

func (e *MaterialTypeGatewayImpl) DeleteMaterialType(id string) error {
	return e.materialTypeRepository.DeleteMaterialType(id)
}

func (e *MaterialTypeGatewayImpl) ToBusinessEntity(materialTypeDB gateway_entities.MaterialType) entities.MaterialType {
	materialType := entities.MaterialType{
		EntitiesBase: core.EntitiesBase{
			ID: materialTypeDB.ID,
		},
		Name:              materialTypeDB.Name,
		Description:       materialTypeDB.Description,
		UnitOfMeasurement: enums.StringToUnitOfMeasurementEnum(materialTypeDB.UnitOfMeasurement),
	}
	return materialType
}

func (e *MaterialTypeGatewayImpl) ToServiceEntity(materialType entities.MaterialType) gateway_entities.MaterialType {
	materialTypeDB := gateway_entities.MaterialType{
		Name:              materialType.Name,
		Description:       materialType.Description,
		UnitOfMeasurement: materialType.UnitOfMeasurement.String(),
	}
	return materialTypeDB
}
