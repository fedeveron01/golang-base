package material_type_handler

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
)

func ToMaterialTypesResponse(materialType []entities.MaterialType, language string) []MaterialTypeResponse {
	var materialTypeResponses []MaterialTypeResponse
	for _, materialType := range materialType {
		materialTypeResponses = append(materialTypeResponses, MaterialTypeResponse{
			Id:                      float64(materialType.ID),
			Name:                    materialType.Name,
			Description:             materialType.Description,
			UnitOfMeasurement:       materialType.UnitOfMeasurement.String(language),
			UnitOfMeasurementSymbol: enums.GetSymbolByUnitOfMeasurementEnum(materialType.UnitOfMeasurement),
		})
	}
	return materialTypeResponses
}

func ToUnitsOfMeasurementResponse(unitsOfMeasurement []string) []UnitOfMeasurementResponse {
	var unitsOfMeasurementResponses []UnitOfMeasurementResponse
	for _, unitOfMeasurement := range unitsOfMeasurement {
		unitsOfMeasurementResponses = append(unitsOfMeasurementResponses, UnitOfMeasurementResponse{
			Name:   unitOfMeasurement,
			Symbol: enums.GetSymbolByUnitOfMeasurementEnum(enums.StringToUnitOfMeasurementEnum(unitOfMeasurement)),
		})
	}
	return unitsOfMeasurementResponses
}
