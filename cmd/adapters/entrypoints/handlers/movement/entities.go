package movement

//import material entity from core

type MovementRequest struct {
	Type               string                  `json:"type"`
	Description        string                  `json:"description" optional:"true"`
	Details            []MovementDetailRequest `json:"details"`
	EmployeeID         uint                    `json:"employeeId"`
	IsMaterialMovement bool                    `json:"isMaterialMovement"`
}

type MovementDetailRequest struct {
	ProductVariationID uint    `json:"productVariationId" optional:"true"`
	MaterialID         uint    `json:"materialId" optional:"true"`
	Quantity           float64 `json:"quantity"`
	Price              float64 `json:"price"`
}

type MovementResponse struct {
	ID                 uint                     `json:"id"`
	Number             int                      `json:"number"`
	Type               string                   `json:"type"`
	Total              float64                  `json:"total"`
	DateTime           string                   `json:"dateTime"`
	Description        string                   `json:"description"`
	MovementDetails    []MovementDetailResponse `json:"movementDetails"`
	IsMaterialMovement bool                     `json:"isMaterialMovement"`
}

type MovementDetailResponse struct {
	ID                 uint    `json:"id"`
	ProductVariationID uint    `json:"productVariationId"`
	MaterialID         uint    `json:"materialId"`
	Quantity           float64 `json:"quantity"`
	Price              float64 `json:"price"`
}

type ProductVariationResponse struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"productId"`
	Number    int     `json:"number"`
	Stock     float64 `json:"stock"`
}

type MovementRequestByType struct {
	IsMaterialMovement bool `json:"isMaterialMovement"`
}
