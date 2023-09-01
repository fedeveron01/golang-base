package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type MaterialGateway interface {
	CreateMaterial(material entities.Material) error
	FindAll() ([]entities.Material, error)
	UpdateMaterial(material entities.Material) error
	DeleteMaterial(id string) error
}

type MaterialGatewayImpl struct {
	materialRepository repositories.MaterialRepository
}

func NewMaterialGateway(materialRepository repositories.MaterialRepository) *MaterialGatewayImpl {
	return &MaterialGatewayImpl{
		materialRepository: materialRepository,
	}
}

func (i *MaterialGatewayImpl) CreateMaterial(material entities.Material) error {

	materialTypeDB := gateway_entities.MaterialType{
		Name: material.MaterialType.Name,
	}
	measurementUnitDB := gateway_entities.MeasurementUnit{
		Name: material.MeasurementUnit.Name,
	}
	materialDB := gateway_entities.Material{
		Name:            material.Name,
		Description:     material.Description,
		Price:           material.Price,
		Stock:           material.Stock,
		MaterialType:    materialTypeDB,
		MeasurementUnit: measurementUnitDB,
	}
	err := i.materialRepository.CreateMaterial(materialDB)
	if err != nil {
		return err
	}
	return nil
}

func (i *MaterialGatewayImpl) FindAll() ([]entities.Material, error) {
	materialsDB, err := i.materialRepository.FindAll()
	if err != nil {
		return nil, err
	}
	materials := make([]entities.Material, len(materialsDB))
	for i, materialDB := range materialsDB {
		materials[i] = entities.Material{
			EntitiesBase: core.EntitiesBase{
				ID: materialDB.ID,
			},
			Name:            materialDB.Name,
			Description:     materialDB.Description,
			Price:           materialDB.Price,
			Stock:           materialDB.Stock,
			MaterialType:    entities.MaterialType{Name: materialDB.MaterialType.Name},
			MeasurementUnit: entities.MeasurementUnit{Name: materialDB.MeasurementUnit.Name},
		}

	}
	return materials, err
}

func (i *MaterialGatewayImpl) UpdateMaterial(material entities.Material) error {
	materialTypeDB := gateway_entities.MaterialType{
		Name: material.MaterialType.Name,
	}
	measurementUnitDB := gateway_entities.MeasurementUnit{
		Name: material.MeasurementUnit.Name,
	}
	materialDB := gateway_entities.Material{
		Name:            material.Name,
		Description:     material.Description,
		Price:           material.Price,
		Stock:           material.Stock,
		MaterialType:    materialTypeDB,
		MeasurementUnit: measurementUnitDB,
	}
	return i.materialRepository.UpdateMaterial(materialDB)
}

func (i *MaterialGatewayImpl) DeleteMaterial(id string) error {
	return i.materialRepository.DeleteMaterial(id)
}
