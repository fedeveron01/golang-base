package product_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
)

type ProductUseCase interface {
	FindAll() ([]entities.Product, error)
	FindById(id uint) (entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	UpdateProduct(product entities.Product) (entities.Product, error)
	DeleteProduct(id uint) error
	AssignMaterialsToProduct(productId float64, materialsProduct []entities.MaterialProduct) ([]entities.MaterialProduct, error)
}

type ProductGateway interface {
	FindAll() ([]entities.Product, error)
	FindById(id uint) *entities.Product
	FindByName(name string) *entities.Product
	CreateProduct(productType entities.Product) (entities.Product, error)
	UpdateProduct(productType entities.Product) (entities.Product, error)
	DeleteProduct(id uint) error
	UpdateMaterialProducts(productId uint, materialsProduct []entities.MaterialProduct) ([]entities.MaterialProduct, error)
}

type MaterialGateway interface {
	FindById(id uint) *entities.Material
}

type Implementation struct {
	productGateway  ProductGateway
	materialGateway MaterialGateway
}

func NewProductUsecase(productGateway ProductGateway, materialGateway MaterialGateway) *Implementation {
	return &Implementation{
		productGateway:  productGateway,
		materialGateway: materialGateway,
	}
}

func (i *Implementation) FindAll() ([]entities.Product, error) {
	materials, err := i.productGateway.FindAll()
	if err != nil {
		return nil, err
	}
	return materials, nil
}

func (i *Implementation) FindById(id uint) (entities.Product, error) {
	product := i.productGateway.FindById(id)
	if product == nil {
		return entities.Product{}, core_errors.NewNotFoundError("product not found")
	}
	return *product, nil
}

func (i *Implementation) CreateProduct(product entities.Product) (entities.Product, error) {
	if product.Name == "" {
		return entities.Product{}, core_errors.NewBadRequestError("product name is required")
	}
	if len(product.Name) < 2 {
		return entities.Product{}, core_errors.NewBadRequestError("product name must be at least 2 characters")
	}
	if len(product.Name) > 20 {
		return entities.Product{}, core_errors.NewBadRequestError("product name must be at most 20 characters")
	}
	repeated := i.productGateway.FindByName(product.Name)
	if repeated != nil {
		return entities.Product{}, core_errors.NewInternalServerError("product already exists")
	}
	var err error
	product, err = i.productGateway.CreateProduct(product)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (i *Implementation) UpdateProduct(product entities.Product) (entities.Product, error) {
	if product.ID <= 0 {
		return entities.Product{}, core_errors.NewInternalServerError("product id is required")
	}

	found := i.productGateway.FindById(product.ID)
	if found == nil && found.ID != product.ID {
		return entities.Product{}, core_errors.NewInternalServerError("product not exist")
	}

	repeated := i.productGateway.FindByName(product.Name)
	if repeated != nil && repeated.ID != product.ID {
		return entities.Product{}, core_errors.NewInternalServerError("product already exists")
	}

	if product.Name == "" {
		return entities.Product{}, core_errors.NewBadRequestError("product name is required")
	}

	if len(product.Name) < 2 {
		return entities.Product{}, core_errors.NewBadRequestError("product name must be at least 2 characters")
	}
	if len(product.Name) > 20 {
		return entities.Product{}, core_errors.NewBadRequestError("product name must be less than 20 characters")
	}

	product, err := i.productGateway.UpdateProduct(product)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil

}

func (i *Implementation) DeleteProduct(id uint) error {
	err := i.productGateway.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

func (i *Implementation) AssignMaterialsToProduct(productId float64, materialsProduct []entities.MaterialProduct) ([]entities.MaterialProduct, error) {
	product := i.productGateway.FindById(uint(productId))
	if product == nil {
		return nil, core_errors.NewNotFoundError("product not found")
	}
	uniqueMaterials := make(map[uint]entities.MaterialProduct)

	for _, materialProduct := range materialsProduct {

		if materialProduct.Material.ID <= 0 {
			return nil, core_errors.NewBadRequestError("material id is required")
		}
		material := i.materialGateway.FindById(materialProduct.Material.ID)
		if material == nil {
			return nil, core_errors.NewNotFoundError("material not found")
		}
		if materialProduct.Quantity <= 0 {
			return nil, core_errors.NewBadRequestError("material quantity is required")
		}
		if uniqueMaterials[materialProduct.Material.ID].Material.ID > 0 {
			return nil, core_errors.NewBadRequestError("material duplicated")
		}
		uniqueMaterials[materialProduct.Material.ID] = materialProduct
	}

	materialsProduct, err := i.productGateway.UpdateMaterialProducts(uint(productId), materialsProduct)
	return materialsProduct, err
}
