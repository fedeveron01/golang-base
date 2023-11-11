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
		MovementDetail: toServiceMovementDetails(movement.MovementDetail, movementID),
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
		MovementDetail: toBusinessMovementDetails(movement.MovementDetail),
	}
}
