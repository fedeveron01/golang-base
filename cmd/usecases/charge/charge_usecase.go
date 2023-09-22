package charge_usecase

import (
	"errors"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

type ChargeUseCase interface {
	CreateCharge(charge entities.Charge) error
}

type ChargeGateway interface {
	CreateCharge(charge entities.Charge) (entities.Charge, error)
	FindByName(name string) (uint, error)
}

type Implementation struct {
	chargeGateway ChargeGateway
}

func NewChargeUsecase(chargeGateway ChargeGateway) *Implementation {
	return &Implementation{
		chargeGateway: chargeGateway,
	}
}

func (i *Implementation) CreateCharge(charge entities.Charge) error {
	if charge.Name == "" {
		return errors.New("name is required")
	}
	if len(charge.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	repeated, err := i.chargeGateway.FindByName(charge.Name)
	if repeated != 0 {
		return errors.New("name already exists")
	}

	_, err = i.chargeGateway.CreateCharge(charge)

	return err
}
