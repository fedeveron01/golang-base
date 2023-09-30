package material_handler

import (
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

func ToMaterialEntity(request MaterialRequest) entities.Material {
	return entities.Material{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		MaterialType: entities.MaterialType{
			EntitiesBase: core.EntitiesBase{
				ID: request.MaterialTypeId,
			},
		},
	}

}

func ToMaterialResponse(material entities.Material) MaterialResponse {
	return MaterialResponse{
		ID:             material.ID,
		Name:           material.Name,
		Description:    material.Description,
		Price:          material.Price,
		Stock:          material.Stock,
		MaterialType:   material.MaterialType.Name,
		MaterialTypeId: material.MaterialType.ID,
	}
}
