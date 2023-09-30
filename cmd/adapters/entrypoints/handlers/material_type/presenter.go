package material_type_handler

import "github.com/fedeveron01/golang-base/cmd/core/entities"

func ToMaterialTypesResponse(materialType []entities.MaterialType) []MaterialTypeResponse {
	var materialTypeResponses []MaterialTypeResponse
	for _, materialType := range materialType {
		materialTypeResponses = append(materialTypeResponses, MaterialTypeResponse{
			Id:   float64(materialType.ID),
			Name: materialType.Name,
		})
	}
	return materialTypeResponses
}
