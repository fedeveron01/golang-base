package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
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

func (i *MovementGatewayImpl) ToServiceEntity(material entities.Movement) gateway_entities.Movement {
	return gateway_entities.Movement{
		ID:        material.ID,
		Type:      material.Type,
		CreatedAt: material.CreatedAt,
		UpdatedAt: material.UpdatedAt,
		DeletedAt: material.DeletedAt,
	}
}
