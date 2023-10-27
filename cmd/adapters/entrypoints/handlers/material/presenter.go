package material_handler

import (
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
)

func ToMaterialResponse(material entities.Material, language string) MaterialResponse {
	return MaterialResponse{
		Id:              float64(material.ID),
		Name:            material.Name,
		Description:     material.Description,
		Price:           material.Price,
		Stock:           material.Stock,
		RepositionPoint: material.RepositionPoint,
		MaterialType: MaterialTypeResponse{
			Id:   float64(material.MaterialType.ID),
			Name: material.MaterialType.Name,
			UnitOfMeasurement: UnitOfMeasurementResponse{
				Name:   material.MaterialType.UnitOfMeasurement.String(language),
				Symbol: enums.GetSymbolByUnitOfMeasurementEnum(material.MaterialType.UnitOfMeasurement),
			},
		},
	}
}

func ToMaterialsResponse(materials []entities.Material, language string) []MaterialResponse {
	var materialResponses []MaterialResponse
	for _, material := range materials {
		materialResponses = append(materialResponses, MaterialResponse{
			float64(material.ID),
			material.Name,
			material.Description,
			MaterialTypeResponse{
				Id:   float64(material.MaterialType.ID),
				Name: material.MaterialType.Name,
				UnitOfMeasurement: UnitOfMeasurementResponse{
					Name:   material.MaterialType.UnitOfMeasurement.String(language),
					Symbol: enums.GetSymbolByUnitOfMeasurementEnum(material.MaterialType.UnitOfMeasurement),
				},
			},
			material.Price,
			material.Stock,
			material.RepositionPoint,
		})
	}
	return materialResponses
}

func ToMaterialEntity(request MaterialRequest) entities.Material {
	return entities.Material{
		EntitiesBase: core.EntitiesBase{
			ID: uint(request.Id),
		},
		Name:            request.Name,
		Description:     request.Description,
		Price:           request.Price,
		Stock:           request.Stock,
		RepositionPoint: request.RepositionPoint,
		MaterialType: entities.MaterialType{
			EntitiesBase: core.EntitiesBase{
				ID: request.MaterialTypeId,
			},
		},
	}
}
