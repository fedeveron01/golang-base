package repositories

import (
	"errors"
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"gorm.io/gorm"
)

type ChargeRepository struct {
	db *gorm.DB
}

func NewChargeRepository(database *gorm.DB) *ChargeRepository {
	return &ChargeRepository{
		db: database,
	}
}

func (r *ChargeRepository) FindByName(name string) (gateway_entities.Charge, error) {
	var charge gateway_entities.Charge
	res := r.db.Where("name = ?", name).Find(&charge)
	if res.Error != nil {
		return gateway_entities.Charge{}, res.Error
	}
	if res.RowsAffected == 0 {
		return gateway_entities.Charge{}, errors.New("charge not found")
	}
	return charge, nil
}

func (r *ChargeRepository) FindById(id uint) (gateway_entities.Charge, error) {
	var charge gateway_entities.Charge
	res := r.db.Where("id = ?", id).Find(&charge)
	if res.Error != nil {
		return gateway_entities.Charge{}, res.Error
	}
	if res.RowsAffected == 0 {
		return gateway_entities.Charge{}, errors.New("charge not found")
	}
	return charge, nil
}

func (r *ChargeRepository) CreateCharge(charge gateway_entities.Charge) (gateway_entities.Charge, error) {
	id := r.db.Create(&charge)
	if id.Error != nil {
		return gateway_entities.Charge{}, id.Error
	}
	if id.RowsAffected == 0 {
		return gateway_entities.Charge{}, errors.New("charge not created")
	}
	return charge, nil
}
