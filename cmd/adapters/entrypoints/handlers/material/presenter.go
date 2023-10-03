package material_handler

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
)

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
