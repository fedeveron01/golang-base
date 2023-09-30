package material_type_handler

import "github.com/fedeveron01/golang-base/cmd/core/entities"

func ToMaterialTypesResponse(materialType []entities.MaterialType, language string) []MaterialTypeResponse {
	var materialTypeResponses []MaterialTypeResponse
	for _, materialType := range materialType {
		materialTypeResponses = append(materialTypeResponses, MaterialTypeResponse{
			Id:                float64(materialType.ID),
			Name:              materialType.Name,
			UnitOfMeasurement: materialType.UnitOfMeasurement.String(language),
		})
	}
	return materialTypeResponses
}

func ToUnitsOfMeasurementResponse(unitsOfMeasurement []string) []UnitOfMeasurementResponse {
	var unitsOfMeasurementResponses []UnitOfMeasurementResponse
	for _, unitOfMeasurement := range unitsOfMeasurement {
		unitsOfMeasurementResponses = append(unitsOfMeasurementResponses, UnitOfMeasurementResponse{
			Name: unitOfMeasurement,
		})
	}
	return unitsOfMeasurementResponses
}
