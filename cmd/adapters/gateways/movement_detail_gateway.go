package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"gorm.io/gorm"
)

type MovementDetailRepository interface {
	CreateMovementDetailsTransaction(movementDetails []gateway_entities.MovementDetail) ([]gateway_entities.MovementDetail, error)
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

func (i *MovementDetailGatewayImpl) CreateMovementDetailsTransaction(movementDetails []entities.MovementDetail, movementID uint) ([]entities.MovementDetail, error) {
	movementDetailsDB := toServiceMovementDetails(movementDetails, movementID)
	movementDetailsDBCreated, err := i.movementDetailRepository.CreateMovementDetailsTransaction(movementDetailsDB)
	if err != nil {
		return nil, err
	}
	movementDetailsCreated := toBusinessMovementDetails(movementDetailsDBCreated)
	return movementDetailsCreated, nil
}

func toServiceMovementDetails(movementDetails []entities.MovementDetail, movementID uint) []gateway_entities.MovementDetail {
	movementDetailsDB := make([]gateway_entities.MovementDetail, len(movementDetails))
	for index, movementDetail := range movementDetails {
		movementDetailsDB[index] = toServiceMovementDetail(movementDetail, movementID)
	}
	return movementDetailsDB
}

func toServiceMovementDetail(movementDetail entities.MovementDetail, movementID uint) gateway_entities.MovementDetail {
	return gateway_entities.MovementDetail{
		Model: gorm.Model{
			ID: movementDetail.ID,
		},
		MaterialId: movementDetail.Material.ID,
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
		ProductVariationId: movementDetail.ProductVariation.ID,
		Quantity:           movementDetail.Quantity,
		MovementId:         movementID,
	}
}

func toBusinessMovementDetails(movementDetails []gateway_entities.MovementDetail) []entities.MovementDetail {
	movementDetailsBusiness := make([]entities.MovementDetail, len(movementDetails))
	for index, movementDetail := range movementDetails {
		movementDetailsBusiness[index] = toBusinessMovementDetail(movementDetail)
	}
	return movementDetailsBusiness
}

func toBusinessMovementDetail(movementDetail gateway_entities.MovementDetail) entities.MovementDetail {
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
