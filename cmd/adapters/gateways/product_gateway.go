package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
	"github.com/fedeveron01/golang-base/cmd/repositories"
	"gorm.io/gorm"
)

type ProductGatewayImpl struct {
	productRepository repositories.ProductRepository
}

func NewProductGateway(productRepository repositories.ProductRepository) *ProductGatewayImpl {
	return &ProductGatewayImpl{
		productRepository: productRepository,
	}
}

func (e *ProductGatewayImpl) FindAll() ([]entities.Product, error) {
	productsDB, err := e.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	products := make([]entities.Product, len(productsDB))
	for i, productDB := range productsDB {
		products[i] = e.ToBusinessEntity(productDB)
	}
	return products, err
}

func (e *ProductGatewayImpl) FindById(id uint) *entities.Product {
	productDB := e.productRepository.FindById(id)
	if productDB == nil {
		return nil
	}
	product := e.ToBusinessEntity(*productDB)
	product.MaterialProduct = e.MaterialsProductToBusinessEntity(productDB.MaterialProduct)
	return &product
}

func (e *ProductGatewayImpl) FindByName(name string) *entities.Product {
	productDB := e.productRepository.FindByName(name)
	if productDB == nil {
		return nil
	}
	product := e.ToBusinessEntity(*productDB)
	return &product
}

func (e *ProductGatewayImpl) CreateProduct(product entities.Product) (entities.Product, error) {
	productDB := e.ToServiceEntity(product)
	created, err := e.productRepository.CreateProduct(productDB)
	if err != nil {
		return entities.Product{}, err
	}
	product = e.ToBusinessEntity(created)
	return product, nil
}

func (e *ProductGatewayImpl) UpdateProduct(product entities.Product) (entities.Product, error) {

	productDB := e.ToServiceEntity(product)
	var err error
	productDB, err = e.productRepository.UpdateProduct(productDB)
	if err != nil {
		return entities.Product{}, err
	}
	product = e.ToBusinessEntity(productDB)
	return product, nil

}

func (e *ProductGatewayImpl) DeleteProduct(id uint) error {
	return e.productRepository.DeleteProduct(id)
}

func (e *ProductGatewayImpl) UpdateMaterialProducts(productId uint, materialsProduct []entities.MaterialProduct) ([]entities.MaterialProduct, error) {
	materialsProductDB := make([]gateway_entities.MaterialProduct, len(materialsProduct))
	for i, materialProduct := range materialsProduct {
		materialsProductDB[i] = gateway_entities.MaterialProduct{
			ProductId:  productId,
			MaterialId: materialProduct.Material.ID,
			Quantity:   materialProduct.Quantity,
		}
	}
	materialsProductDB, err := e.productRepository.UpdateMaterialProducts(int64(productId), materialsProductDB)
	if err != nil {
		return nil, err
	}
	materialsProduct = e.MaterialsProductToBusinessEntity(materialsProductDB)
	return materialsProduct, nil
}

func (e *ProductGatewayImpl) ToBusinessEntity(product gateway_entities.Product) entities.Product {
	productBusiness := entities.Product{
		EntitiesBase: core.EntitiesBase{
			ID: product.ID,
		},
		Name:        product.Name,
		Description: product.Description,
		Color:       product.Color,
		Size:        product.Size,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
		Stock:       product.Stock,
	}
	return productBusiness
}

func (e *ProductGatewayImpl) ToServiceEntity(product entities.Product) gateway_entities.Product {
	productDB := gateway_entities.Product{
		Model: gorm.Model{
			ID: product.ID,
		},
		Name:        product.Name,
		Description: product.Description,
		Color:       product.Color,
		Size:        product.Size,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
		Stock:       product.Stock,
	}
	return productDB
}

func (e *ProductGatewayImpl) MaterialsProductToBusinessEntity(materialProducts []gateway_entities.MaterialProduct) []entities.MaterialProduct {
	materialProductsBusiness := make([]entities.MaterialProduct, len(materialProducts))
	for i, materialProduct := range materialProducts {
		materialProductsBusiness[i] = entities.MaterialProduct{
			EntitiesBase: core.EntitiesBase{
				ID: materialProduct.ID,
			},
			Material: entities.Material{
				EntitiesBase: core.EntitiesBase{
					ID: materialProduct.Material.ID,
				},
				Name: materialProduct.Material.Name,
				MaterialType: entities.MaterialType{
					EntitiesBase: core.EntitiesBase{
						ID: materialProduct.Material.MaterialType.ID,
					},
					Name:              materialProduct.Material.MaterialType.Name,
					UnitOfMeasurement: enums.StringToUnitOfMeasurementEnum(materialProduct.Material.MaterialType.UnitOfMeasurement),
				},
			},

			Quantity: materialProduct.Quantity,
		}
	}
	return materialProductsBusiness
}
