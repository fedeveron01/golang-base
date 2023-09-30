package material_handler

type MaterialResponse struct {
	Id              float64 `json:"id"`
	Name            string  `json:"name"`
	MaterialType    string  `json:"materialType"`
	Price           float64 `json:"price"`
	Stock           int     `json:"stock"`
	RepositionPoint int     `json:"repositionPoint"`
}
