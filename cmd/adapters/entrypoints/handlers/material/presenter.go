package material_handler

import "github.com/fedeveron01/golang-base/cmd/core/entities"

func ToMaterialsResponse(materials []entities.Material) []MaterialResponse {
	var materialResponses []MaterialResponse
	for _, material := range materials {
		materialResponses = append(materialResponses, MaterialResponse{
			Id:              float64(material.ID),
			Name:            material.Name,
			MaterialType:    material.MaterialType.Name,
			Price:           material.Price,
			Stock:           material.Stock,
			RepositionPoint: material.RepositionPoint,
		})
	}
	return materialResponses
}
