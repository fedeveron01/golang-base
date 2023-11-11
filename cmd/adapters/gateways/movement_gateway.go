package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
)

type MovementRepository interface {
	CreateMovement(movement gateway_entities.Movement) (gateway_entities.Movement, error)
}

type MovementGateway interface {
	Create(movement entities.Movement) (entities.Movement, error)
}

type MovementGatewayImpl struct {
	movementRepository MovementRepository
}

func NewMovementGateway(movementRepository MovementRepository) *MovementGatewayImpl {
	return &MovementGatewayImpl{
		movementRepository: movementRepository,
	}
}
func (i *MovementGatewayImpl) Create(movement entities.Movement, employeeID uint) (entities.Movement, error) {
	movementDB := i.ToServiceEntity(movement, movement.ID)
	movementDB.EmployeeId = employeeID
	movementDB.MovementDetail = nil
	movementDBCreated, err := i.movementRepository.CreateMovement(movementDB)
	if err != nil {
		return entities.Movement{}, err
	}
	movementCreated := i.ToBusinessEntity(movementDBCreated)
	return movementCreated, nil
}

func (i *MovementGatewayImpl) ToServiceEntity(movement entities.Movement, movementID uint) gateway_entities.Movement {
	return gateway_entities.Movement{
		Number:         float64(movement.Number),
		Type:           movement.Type,
		Total:          movement.Total,
		DateTime:       movement.DateTime,
		Description:    movement.Description,
		MovementDetail: i.toServiceMovementDetails(movement.MovementDetail, movementID),
	}
}

func (i *MovementGatewayImpl) ToBusinessEntity(movement gateway_entities.Movement) entities.Movement {
	return entities.Movement{
		EntitiesBase: core.EntitiesBase{
			ID: movement.ID,
		},
		Number:         int(movement.Number),
		Type:           movement.Type,
		Total:          movement.Total,
		DateTime:       movement.DateTime,
		Description:    movement.Description,
		MovementDetail: i.toBusinessMovementDetails(movement.MovementDetail),
	}
}

func (i *MovementGatewayImpl) toServiceMovementDetails(movementDetails []entities.MovementDetail, movementID uint) []gateway_entities.MovementDetail {
	movementDetailsDB := make([]gateway_entities.MovementDetail, len(movementDetails))
	for index, movementDetail := range movementDetails {
		movementDetailsDB[index] = i.toServiceMovementDetail(movementDetail, movementID)
	}
	return movementDetailsDB
}

func (i *MovementGatewayImpl) toServiceMovementDetail(movementDetail entities.MovementDetail, movementID uint) gateway_entities.MovementDetail {
	var materialID uint
	if movementDetail.Material != nil {
		materialID = movementDetail.Material.ID
	}
	var productVariationID uint
	if movementDetail.ProductVariation != nil {
		productVariationID = movementDetail.ProductVariation.ID
	}
	return gateway_entities.MovementDetail{
		MaterialId:         &materialID,
		ProductVariationId: &productVariationID,
		Quantity:           movementDetail.Quantity,
		MovementId:         movementID,
	}
}

func (i *MovementGatewayImpl) toBusinessMovementDetails(movementDetails []gateway_entities.MovementDetail) []entities.MovementDetail {
	movementDetailsBusiness := make([]entities.MovementDetail, len(movementDetails))
	for index, movementDetail := range movementDetails {
		movementDetailsBusiness[index] = i.toBusinessMovementDetail(movementDetail)
	}
	return movementDetailsBusiness
}

func (i *MovementGatewayImpl) toBusinessMovementDetail(movementDetail gateway_entities.MovementDetail) entities.MovementDetail {
	var material *entities.Material
	var productVariation *entities.ProductVariation
	if movementDetail.Material != nil && movementDetail.Material.ID != 0 {
		material = &entities.Material{
			EntitiesBase: core.EntitiesBase{
				ID: movementDetail.Material.ID,
			},
		}
	}
	if movementDetail.ProductVariation != nil && movementDetail.ProductVariation.ID != 0 {
		productVariation = &entities.ProductVariation{
			EntitiesBase: core.EntitiesBase{
				ID: movementDetail.ProductVariation.ID,
			},
		}
	}

	return entities.MovementDetail{
		Material:         material,
		ProductVariation: productVariation,
		Quantity:         movementDetail.Quantity,
	}
}