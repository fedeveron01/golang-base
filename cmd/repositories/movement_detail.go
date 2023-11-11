package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"gorm.io/gorm"
)

type MovementDetailRepository struct {
	db *gorm.DB
}

func NewMovementDetailRepository(database *gorm.DB) *MovementDetailRepository {
	return &MovementDetailRepository{
		db: database,
	}
}

func (r *MovementDetailRepository) CreateMovementDetailsTransaction(movementDetails []gateway_entities.MovementDetail, movement gateway_entities.Movement) ([]gateway_entities.MovementDetail, *gateway_entities.Movement, error) {
	tx := r.db.Begin()
	tx.Create(&movement)
	for _, movementDetail := range movementDetails {
		movementDetail.MovementId = movement.ID
		if movementDetail.Material != nil && movementDetail.Material.ID != 0 {
			res := tx.Save(&movementDetail.Material)
			if res.Error != nil {
				tx.Rollback()
				return nil, nil, res.Error
			}
		}
		movementDetail.ProductVariationId = nil
		res := tx.Create(&movementDetail)
		if res.Error != nil {
			tx.Rollback()
			return nil, nil, res.Error
		}
	}
	tx.Commit()
	return movementDetails, &movement, nil
}
