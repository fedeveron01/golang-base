package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	"gorm.io/gorm"
)

type MaterialTypeRepository struct {
	db *gorm.DB
}

func NewMaterialTypeRepository(database *gorm.DB) *MaterialTypeRepository {
	return &MaterialTypeRepository{
		db: database,
	}
}

func (r *MaterialTypeRepository) CreateMaterialType(materialType gateway_entities.MaterialType) error {
	id := r.db.Create(&materialType)
	if id.Error != nil {
		return id.Error
	}
	return nil
}

func (r *MaterialTypeRepository) FindAll() ([]gateway_entities.MaterialType, error) {
	var materialTypes []gateway_entities.MaterialType
	r.db.Find(&materialTypes)
	return materialTypes, nil
}

func (r *MaterialTypeRepository) FindByName(name string) *gateway_entities.MaterialType {
	var materialType gateway_entities.MaterialType
	r.db.Where("name = ?", name).First(&materialType)
	if materialType.ID == 0 {
		return nil
	}
	return &materialType
}

func (r *MaterialTypeRepository) UpdateMaterialType(materialType gateway_entities.MaterialType) (gateway_entities.MaterialType, error) {
	res := r.db.Save(&materialType)
	if res.Error != nil {
		return gateway_entities.MaterialType{}, res.Error
	}
	if res.RowsAffected == 0 {
		return gateway_entities.MaterialType{}, core_errors.NewInternalServerError("materialType update failed")
	}
	return materialType, nil
}

func (r *MaterialTypeRepository) DeleteMaterialType(id string) error {
	r.db.Delete(&gateway_entities.MaterialType{}, id)
	return nil
}
