package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type ChargeGatewayImpl struct {
	chargeRepository repositories.ChargeRepository
}

func NewChargeGateway(chargeRepository repositories.ChargeRepository) *ChargeGatewayImpl {
	return &ChargeGatewayImpl{
		chargeRepository: chargeRepository,
	}
}

func (c ChargeGatewayImpl) FindByName(name string) (uint, error) {
	charge, err := c.chargeRepository.FindByName(name)
	if err != nil {
		return 0, err
	}
	return charge.ID, nil
}

func (c ChargeGatewayImpl) CreateCharge(charge entities.Charge) (entities.Charge, error) {
	chargeDB := gateway_entities.Charge{
		Name: charge.Name,
	}
	created, err := c.chargeRepository.CreateCharge(chargeDB)
	if err != nil {
		return entities.Charge{}, err
	}
	charge = entities.Charge{
		EntitiesBase: core.EntitiesBase{
			ID: created.ID,
		},
		Name: created.Name,
	}
	return charge, nil
}