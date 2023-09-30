package material_handler

type MaterialResponse struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Stock           int     `json:"stock"`
	RepositionPoint int     `json:"repositionPoint"`
	MaterialType    string  `json:"materialType"`
	MaterialTypeId  uint    `json:"materialTypeId"`
}

type MaterialRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Stock           int     `json:"stock"`
	RepositionPoint int     `json:"repositionPoint"`
	MaterialTypeId  uint    `json:"materialTypeId"`
}
