package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"gorm.io/gorm"
)

type MovementRepository struct {
	db *gorm.DB
}

func NewMovementRepository(database *gorm.DB) *MovementRepository {
	return &MovementRepository{
		db: database,
	}
}

func (r *MovementRepository) CreateMovement(movement gateway_entities.Movement) (gateway_entities.Movement, error) {
	id := r.db.Create(&movement)
	if id.Error != nil {
		return gateway_entities.Movement{}, id.Error
	}
	return movement, nil
}
