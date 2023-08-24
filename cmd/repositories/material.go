package repositories

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
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

func (r *MaterialRepository) CreateMaterial(material entities.Material) error {
	id := r.db.Create(&material)
	if id.Error != nil {
		return id.Error
	}
	return nil
}

func (r *MaterialRepository) FindAll() ([]entities.Material, error) {
	var materials []entities.Material
	r.db.Find(&materials)
	return materials, nil
}

func (r *MaterialRepository) UpdateMaterial(material entities.Material) error {
	r.db.Save(&material)
	return nil
}

func (r *MaterialRepository) DeleteMaterial(id string) error {
	r.db.Delete(&entities.Material{}, id)
	return nil
}
