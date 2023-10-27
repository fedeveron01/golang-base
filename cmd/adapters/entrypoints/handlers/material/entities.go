package material_handler

type MaterialRequest struct {
	Id              float64 `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Stock           float64 `json:"stock"`
	RepositionPoint float64 `json:"repositionPoint"`
	MaterialTypeId  uint    `json:"materialTypeId"`
}

type MaterialResponse struct {
	Id              float64              `json:"id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	MaterialType    MaterialTypeResponse `json:"materialType"`
	Price           float64              `json:"price"`
	Stock           float64              `json:"stock"`
	RepositionPoint float64              `json:"repositionPoint"`
}

type MaterialTypeResponse struct {
	Id                float64                   `json:"id"`
	Name              string                    `json:"name"`
	UnitOfMeasurement UnitOfMeasurementResponse `json:"unitOfMeasurement"`
}

type UnitOfMeasurementResponse struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
