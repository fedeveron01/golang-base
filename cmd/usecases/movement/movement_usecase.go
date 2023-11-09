package movement_usecase

import (
	"errors"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

type MovementUseCase interface {
	Create(movement entities.Movement) (entities.Movement, error)
}

type MovementGateway interface {
	Create(movement entities.Movement) entities.Movement
	CreateMovementDetailsTransaction(movementDetails []entities.MovementDetail) ([]entities.MovementDetail, error)
}

type MaterialGateway interface {
	FindMaterialById(id uint) *entities.Material
}

type MovementUseCaseImpl struct {
	movementGateway MovementGateway
	materialGateway MaterialGateway
}

func NewMovementUseCase(movementGateway MovementGateway, materialGateway MaterialGateway) *MovementUseCaseImpl {
	return &MovementUseCaseImpl{
		movementGateway: movementGateway,
		materialGateway: materialGateway,
	}
}

func (i *MovementUseCaseImpl) updateMaterial(movementDetail *entities.MovementDetail, input bool) error {
	material := i.materialGateway.FindMaterialById(movementDetail.Material.ID)
	if material == nil {
		return errors.New("material not found")
	}

	if input {
		material.Stock += movementDetail.Quantity
	} else {
		if material.Stock-movementDetail.Quantity < 0 {
			return errors.New("insufficient stock")
		}
		material.Stock -= movementDetail.Quantity
	}

	movementDetail.Material = material

	return nil
}

func (i *MovementUseCaseImpl) Create(movement entities.Movement) (entities.Movement, error) {
	if movement.Type == "" {
		return entities.Movement{}, errors.New("type is required")
	}

	if movement.Type == "input" {
		for _, movementDetail := range movement.MovementDetail {
			if movementDetail.Material.ID == 0 && movementDetail.ProductVariation.ID == 0 {
				return entities.Movement{}, errors.New("material or product variation required")
			}
			if movementDetail.Material.ID != 0 {
				err := i.updateMaterial(&movementDetail, true)
				if err != nil {
					return entities.Movement{}, err
				}

			}
			if movementDetail.ProductVariation.ID != 0 {
				panic("implement me please angelo")
			}
		}
	}
	movementCreated := i.movementGateway.Create(movement)

	movementDetailsCreated, err := i.movementGateway.CreateMovementDetailsTransaction(movementCreated.MovementDetail)
	if err != nil {
	}
	movementCreated.MovementDetail = movementDetailsCreated
	return movementCreated, nil

}
