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
	ProductID  uint    `json:"productId" optional:"true"`
	Number     float64 `json:"number" optional:"true"`
	MaterialID uint    `json:"materialId" optional:"true"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
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
	ID                 uint              `json:"id"`
	ProductVariationID uint              `json:"productVariationId"`
	MaterialID         *uint             `json:"materialId"`
	Material           *MaterialResponse `json:"material"`
	Quantity           float64           `json:"quantity"`
	Price              float64           `json:"price"`
}

type MaterialResponse struct {
	ID                      uint    `json:"id"`
	Name                    string  `json:"name"`
	Description             string  `json:"description"`
	MaterialType            string  `json:"materialType"`
	UnitOfMeasurement       string  `json:"unitOfMeasurement"`
	UnitOfMeasurementSymbol string  `json:"unitOfMeasurementSymbol"`
	Stock                   float64 `json:"stock"`
}

type ProductVariationResponse struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"productId"`
	Number    int     `json:"number"`
	Stock     float64 `json:"stock"`
}

type MovementRequestByType struct {
	IsInput bool `json:"isInput"`
}
