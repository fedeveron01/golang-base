package product_handler

import "github.com/fedeveron01/golang-base/cmd/core/entities"

func ToProductsResponse(products []entities.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

func ToProductResponse(product entities.Product) ProductResponse {
	return ProductResponse{
		Id:          float64(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Color:       product.Color,
		Size:        product.Size,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func ToProductWithAssignationsResponse(product entities.Product, language string) ProductWithAssignationsResponse {
	return ProductWithAssignationsResponse{
		Id:           float64(product.ID),
		Name:         product.Name,
		Description:  product.Description,
		Color:        product.Color,
		Size:         product.Size,
		ImageUrl:     product.ImageUrl,
		Price:        product.Price,
		Stock:        product.Stock,
		Assignations: ToAssignationsResponse(product.MaterialProduct, language),
	}
}

func ToAssignationsResponse(materialProducts []entities.MaterialProduct, language string) []AssignationResponse {
	var assignationsResponse []AssignationResponse
	for _, materialProduct := range materialProducts {
		assignationsResponse = append(assignationsResponse, ToAssignationResponse(materialProduct, language))
	}
	return assignationsResponse
}

func ToAssignationResponse(materialProduct entities.MaterialProduct, language string) AssignationResponse {
	return AssignationResponse{
		Id:       float64(materialProduct.ID),
		Quantity: materialProduct.Quantity,
		Material: ToMaterialResponse(materialProduct.Material, language),
	}
}

func ToMaterialResponse(material entities.Material, language string) MaterialResponse {
	return MaterialResponse{
		Id:           float64(material.ID),
		Name:         material.Name,
		MaterialType: ToMaterialTypeResponse(material.MaterialType, language),
	}
}

func ToMaterialTypeResponse(materialType entities.MaterialType, language string) MaterialTypeResponse {
	return MaterialTypeResponse{
		Id:                float64(materialType.ID),
		Name:              materialType.Name,
		UnitOfMeasurement: materialType.UnitOfMeasurement.String(language),
	}
}
