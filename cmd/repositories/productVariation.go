package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"gorm.io/gorm"
)

type ProductVariationRepository struct {
	db *gorm.DB
}

func NewProductVariationRepository(database *gorm.DB) *ProductVariationRepository {
	return &ProductVariationRepository{
		db: database,
	}
}

func (r *ProductVariationRepository) FindById(id uint) *gateway_entities.ProductVariation {
	var productVariation gateway_entities.ProductVariation
	res := r.db.Where("id = ?", id).Find(&productVariation)
	if res.RowsAffected == 0 {
		return nil
	}
	return &productVariation
}
