package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type MovementGateway interface {
	Create(movement entities.Movement) error
}

type MovementGatewayImpl struct {
	materialRepository repositories.MovementRepository
}

func NewMovementGateway(materialRepository repositories.MovementRepository) *MovementGatewayImpl {
	return &MovementGatewayImpl{
		materialRepository: materialRepository,
	}
}

func (i *MovementGatewayImpl) CreateMovement(material entities.Movement) (entities.Movement, error) {

	materialDB := i.ToServiceEntity(material)
	materialDBCreated, err := i.materialRepository.CreateMovement(materialDB)
	if err != nil {
		return entities.Movement{}, err
	}
	materialCreated := i.ToBusinessEntity(materialDBCreated)

	return materialCreated, nil
}

func (i *MovementGatewayImpl) ToServiceEntity(movement entities.Movement) gateway_entities.Movement {
	return gateway_entities.Movement{
		Number:         float64(movement.Number),
		Type:           movement.Type,
		Total:          movement.Total,
		DateTime:       movement.DateTime,
		Description:    movement.Description,
		MovementDetail: toServiceMovementDetail(movement.MovementDetail),
	}
}

func toServiceMovementDetail(movementDetail []entities.MovementDetail) []gateway_entities.MovementDetail {
	var movementDetailDB []gateway_entities.MovementDetail
	for _, v := range movementDetail {
		movementDetailDB = append(movementDetailDB, gateway_entities.MovementDetail{
			Quantity:           v.Quantity,
			Price:              v.Price,
			MaterialId:         v.Material.ID,
			ProductVariationId: v.ProductVariation.ID,
		})
	}
	return movementDetailDB
}

func (i *MovementGatewayImpl) ToBusinessEntity(movement gateway_entities.Movement) entities.Movement {
	return entities.Movement{
		Number:         int(movement.Number),
		Type:           movement.Type,
		Total:          movement.Total,
		DateTime:       movement.DateTime,
		Description:    movement.Description,
		MovementDetail: toBusinessMovementDetail(movement.MovementDetail),
	}
}

func toBusinessMovementDetail(movementDetail []gateway_entities.MovementDetail) []entities.MovementDetail {
	var movementDetailDB []entities.MovementDetail
	for _, v := range movementDetail {
		var material *entities.Material
		var productVariation *entities.ProductVariation
		if v.Material != nil {
			material = &entities.Material{
				EntitiesBase: core.EntitiesBase{
					ID: v.Material.ID,
				},
			}
		} else {
			material = nil
		}
		if v.ProductVariation != nil {
			productVariation = &entities.ProductVariation{
				EntitiesBase: core.EntitiesBase{
					ID: v.ProductVariation.ID,
				},
			}
		} else {
			productVariation = nil
		}

		movementDetailDB = append(movementDetailDB, entities.MovementDetail{
			EntitiesBase: core.EntitiesBase{
				ID: v.ID,
			},
			Quantity:         v.Quantity,
			Price:            v.Price,
			Material:         material,
			ProductVariation: productVariation,
		})
	}
	return movementDetailDB
}
