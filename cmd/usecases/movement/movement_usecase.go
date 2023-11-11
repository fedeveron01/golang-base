package movement_usecase

import (
	"errors"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

type MovementUseCase interface {
	Create(movement entities.Movement, employeeId uint) (entities.Movement, error)
}

type MovementGateway interface {
	Create(movement entities.Movement, employeeID uint) (entities.Movement, error)
}

type MovementDetailGateway interface {
	CreateMovementDetailsTransaction(movementDetails []entities.MovementDetail, movement entities.Movement, employeeID uint) ([]entities.MovementDetail, entities.Movement, error)
}
type MaterialGateway interface {
	FindMaterialById(id uint) *entities.Material
}

type MovementUseCaseImpl struct {
	movementGateway       MovementGateway
	movementDetailGateway MovementDetailGateway
	materialGateway       MaterialGateway
}

func NewMovementUseCase(movementGateway MovementGateway, movementDetailGateway MovementDetailGateway, materialGateway MaterialGateway) *MovementUseCaseImpl {
	return &MovementUseCaseImpl{
		movementGateway:       movementGateway,
		movementDetailGateway: movementDetailGateway,
		materialGateway:       materialGateway,
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
			return errors.New("insufficient stock in material " + material.Name)
		}
		material.Stock -= movementDetail.Quantity
	}

	movementDetail.Material = material

	return nil
}

func (i *MovementUseCaseImpl) Create(movement entities.Movement, employeeID uint) (entities.Movement, error) {
	if movement.Type == "" {
		return entities.Movement{}, errors.New("type is required")
	}
	if employeeID <= 0 {
		return entities.Movement{}, errors.New("employee is required")
	}

	if movement.Type == "input" {
		for index, movementDetail := range movement.MovementDetail {
			if movementDetail.Material.ID == 0 && movementDetail.ProductVariation.ID == 0 {
				return entities.Movement{}, errors.New("material or product variation required")
			}
			if movementDetail.Material.ID != 0 {
				err := i.updateMaterial(&movementDetail, true)
				if err != nil {
					return entities.Movement{}, err
				}
				movement.MovementDetail[index] = movementDetail

			}
			if movementDetail.ProductVariation.ID != 0 {
				panic("implement me please angelo")
			}

		}
	}
	if movement.Type == "output" {
		for index, movementDetail := range movement.MovementDetail {
			if movementDetail.Material.ID == 0 && movementDetail.ProductVariation.ID == 0 {
				return entities.Movement{}, errors.New("material or product variation required")
			}
			if movementDetail.Material.ID != 0 {
				err := i.updateMaterial(&movementDetail, false)
				if err != nil {
					return entities.Movement{}, err
				}
				movement.MovementDetail[index] = movementDetail

			}
			if movementDetail.ProductVariation.ID != 0 {
				panic("implement me please angelo")
			}

		}
	}
	movementDetails := movement.MovementDetail

	movementDetailsCreated, movementCreated, err := i.movementDetailGateway.CreateMovementDetailsTransaction(movementDetails, movement, employeeID)
	if err != nil {
		return entities.Movement{}, err
	}
	movementCreated.MovementDetail = movementDetailsCreated
	return movementCreated, nil

}
