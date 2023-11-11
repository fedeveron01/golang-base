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

func (r *MovementDetailRepository) CreateMovementDetailsTransaction(movementDetails []gateway_entities.MovementDetail) ([]gateway_entities.MovementDetail, error) {
	tx := r.db.Begin()
	for _, movementDetail := range movementDetails {

		if movementDetail.Material.ID != 0 {
			res := tx.Save(&movementDetail.Material)
			if res.Error != nil {
				tx.Rollback()
				return nil, res.Error
			}
		}
		tx.Create(&movementDetail)
	}
	tx.Commit()
	return movementDetails, nil
}
