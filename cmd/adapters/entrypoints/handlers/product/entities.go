package product_handler

type MaterialsProductAssignationRequest struct {
	ProductId    float64              `json:"productId"`
	Assignations []AssignationRequest `json:"assignations"`
}

type AssignationRequest struct {
	MaterialId float64 `json:"materialId"`
	Quantity   float64 `json:"quantity"`
}

type ProductResponse struct {
	Id          float64 `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
	Size        float64 `json:"size"`
	ImageUrl    string  `json:"imageUrl"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type ProductWithAssignationsResponse struct {
	Id           float64               `json:"id"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Color        string                `json:"color"`
	Size         float64               `json:"size"`
	ImageUrl     string                `json:"imageUrl"`
	Price        float64               `json:"price"`
	Stock        int                   `json:"stock"`
	Assignations []AssignationResponse `json:"assignations"`
}

type AssignationResponse struct {
	Quantity float64          `json:"quantity"`
	Material MaterialResponse `json:"material"`
}

type MaterialResponse struct {
	Id           float64              `json:"id"`
	Name         string               `json:"name"`
	MaterialType MaterialTypeResponse `json:"materialType"`
}

type MaterialTypeResponse struct {
	Id                float64 `json:"id"`
	Name              string  `json:"name"`
	UnitOfMeasurement string  `json:"unitOfMeasurement"`
}
