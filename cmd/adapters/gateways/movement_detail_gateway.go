package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"gorm.io/gorm"
)

type MovementDetailRepository interface {
	CreateMovementDetailsTransaction(movementDetails []gateway_entities.MovementDetail, movement gateway_entities.Movement) ([]gateway_entities.MovementDetail, *gateway_entities.Movement, error)
}

type MovementDetailGateway interface {
	CreateMovementDetailsTransaction(movementDetails []entities.MovementDetail) ([]entities.MovementDetail, error)
}

type MovementDetailGatewayImpl struct {
	movementDetailRepository MovementDetailRepository
}

func NewMovementDetailGateway(movementDetailRepository MovementDetailRepository) *MovementDetailGatewayImpl {
	return &MovementDetailGatewayImpl{
		movementDetailRepository: movementDetailRepository,
	}
}

func (i *MovementDetailGatewayImpl) CreateMovementDetailsTransaction(movementDetails []entities.MovementDetail, movement entities.Movement, employeeID uint) ([]entities.MovementDetail, entities.Movement, error) {
	movementDetailsDB := i.toServiceMovementDetails(movementDetails, movement.ID)
	movementDB := i.toServiceMovement(movement, employeeID)

	movementDetailsDBCreated, movementDBCreated, err := i.movementDetailRepository.CreateMovementDetailsTransaction(movementDetailsDB, movementDB)
	if err != nil {
		return nil, entities.Movement{}, err
	}
	movementCreated := i.toBusinessMovement(*movementDBCreated)

	movementDetailsCreated := i.toBusinessMovementDetails(movementDetailsDBCreated)
	return movementDetailsCreated, movementCreated, nil
}

func (i *MovementDetailGatewayImpl) toServiceMovementDetails(movementDetails []entities.MovementDetail, movementID uint) []gateway_entities.MovementDetail {
	movementDetailsDB := make([]gateway_entities.MovementDetail, len(movementDetails))
	for index, movementDetail := range movementDetails {
		movementDetailsDB[index] = i.toServiceMovementDetail(movementDetail, movementID)
	}
	return movementDetailsDB
}

func (i *MovementDetailGatewayImpl) toServiceMovementDetail(movementDetail entities.MovementDetail, movementID uint) gateway_entities.MovementDetail {
	return gateway_entities.MovementDetail{
		Model: gorm.Model{
			ID: movementDetail.ID,
		},
		MaterialId: &movementDetail.Material.ID,
		Material: &gateway_entities.Material{
			Model: gorm.Model{
				ID: movementDetail.Material.ID,
			},
			Name:            movementDetail.Material.Name,
			Description:     movementDetail.Material.Description,
			Price:           movementDetail.Material.Price,
			Stock:           movementDetail.Material.Stock,
			RepositionPoint: movementDetail.Material.RepositionPoint,
			MaterialTypeId:  movementDetail.Material.MaterialType.ID,
		},
		ProductVariationId: &movementDetail.ProductVariation.ID,
		Quantity:           movementDetail.Quantity,
		MovementId:         movementID,
	}
}

func (i *MovementDetailGatewayImpl) toBusinessMovementDetails(movementDetails []gateway_entities.MovementDetail) []entities.MovementDetail {
	movementDetailsBusiness := make([]entities.MovementDetail, len(movementDetails))
	for index, movementDetail := range movementDetails {
		movementDetailsBusiness[index] = i.toBusinessMovementDetail(movementDetail)
	}
	return movementDetailsBusiness
}

func (i *MovementDetailGatewayImpl) toBusinessMovementDetail(movementDetail gateway_entities.MovementDetail) entities.MovementDetail {
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
		EntitiesBase: core.EntitiesBase{
			ID: movementDetail.ID,
		},
		Material:         material,
		ProductVariation: productVariation,
		Quantity:         movementDetail.Quantity,
	}
}

func (i *MovementDetailGatewayImpl) toServiceMovement(movement entities.Movement, employeeID uint) gateway_entities.Movement {
	return gateway_entities.Movement{
		Model: gorm.Model{
			ID: movement.ID,
		},
		Number:             float64(movement.Number),
		Type:               movement.Type,
		Total:              movement.Total,
		DateTime:           movement.DateTime,
		IsMaterialMovement: movement.IsMaterialMovement,
		Description:        movement.Description,
		EmployeeId:         employeeID,
	}
}

func (i *MovementDetailGatewayImpl) toBusinessMovement(movement gateway_entities.Movement) entities.Movement {
	return entities.Movement{
		EntitiesBase: core.EntitiesBase{
			ID: movement.ID,
		},
		Number:             int(movement.Number),
		Type:               movement.Type,
		Total:              movement.Total,
		IsMaterialMovement: movement.IsMaterialMovement,
		DateTime:           movement.DateTime,
		Description:        movement.Description,
	}
}
