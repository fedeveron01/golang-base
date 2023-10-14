package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: database,
	}
}

func (r *ProductRepository) CreateProduct(materialType gateway_entities.Product) (gateway_entities.Product, error) {
	id := r.db.Create(&materialType)
	if id.Error != nil {
		return gateway_entities.Product{}, id.Error
	}
	return materialType, nil
}

func (r *ProductRepository) FindAll() ([]gateway_entities.Product, error) {
	var materialTypes []gateway_entities.Product
	r.db.Find(&materialTypes)
	return materialTypes, nil
}

func (r *ProductRepository) FindByName(name string) *gateway_entities.Product {
	var materialType gateway_entities.Product
	r.db.Where("name = ?", name).First(&materialType)
	if materialType.ID == 0 {
		return nil
	}
	return &materialType
}

func (r *ProductRepository) FindById(id uint) *gateway_entities.Product {
	var materialType gateway_entities.Product
	r.db.InnerJoins("MaterialProduct").InnerJoins("MaterialProduct.MaterialType").
		Find(&materialType, id).First(&materialType)
	if materialType.ID == 0 {
		return nil
	}
	return &materialType
}

func (r *ProductRepository) UpdateProduct(materialType gateway_entities.Product) (gateway_entities.Product, error) {
	res := r.db.Save(&materialType)
	if res.Error != nil {
		return gateway_entities.Product{}, res.Error
	}
	if res.RowsAffected == 0 {
		return gateway_entities.Product{}, core_errors.NewInternalServerError("materialType update failed")
	}
	return materialType, nil
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	r.db.Delete(&gateway_entities.Product{}, id)
	return nil
}
