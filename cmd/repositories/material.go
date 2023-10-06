package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	"gorm.io/gorm"
)

type MaterialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(database *gorm.DB) *MaterialRepository {
	return &MaterialRepository{
		db: database,
	}
}

func (r *MaterialRepository) CreateMaterial(material gateway_entities.Material) (gateway_entities.Material, error) {
	id := r.db.Create(&material)
	if id.Error != nil {
		return gateway_entities.Material{}, id.Error
	}
	var materialType gateway_entities.MaterialType
	r.db.First(&materialType, material.MaterialTypeId)
	material.MaterialType = materialType
	return material, nil
}

func (r *MaterialRepository) FindAll() ([]gateway_entities.Material, error) {
	var materials []gateway_entities.Material
	r.db.InnerJoins("MaterialType").Find(&materials)
	return materials, nil
}

func (r *MaterialRepository) FindByName(name string) *gateway_entities.Material {
	var material gateway_entities.Material
	r.db.Where("name = ?", name).First(&material)
	if material.ID == 0 {
		return nil
	}
	return &material
}

func (r *MaterialRepository) UpdateMaterial(material gateway_entities.Material) error {
	r.db.Save(&material)
	return nil
}

func (r *MaterialRepository) DeleteMaterial(id string) error {
	result := r.db.Delete(&gateway_entities.Material{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return core_errors.NewInternalServerError("Material Id not found")
	}
	return nil
}
