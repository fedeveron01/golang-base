package movement

import (
	"time"

	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

func ToMovement(movementRequest MovementRequest) entities.Movement {
	return entities.Movement{
		Type:           movementRequest.Type,
		DateTime:       time.Now(),
		Description:    movementRequest.Description,
		MovementDetail: ToMovementDetails(movementRequest.Details),
	}
}

func ToMovementDetails(movementDetailsRequest []MovementDetailRequest) []entities.MovementDetail {
	var movementDetails []entities.MovementDetail
	for _, movementDetailRequest := range movementDetailsRequest {
		movementDetails = append(movementDetails, ToMovementDetail(movementDetailRequest))
	}
	return movementDetails
}

func ToMovementDetail(movementDetailRequest MovementDetailRequest) entities.MovementDetail {
	return entities.MovementDetail{
		Quantity: movementDetailRequest.Quantity,
		Price:    movementDetailRequest.Price,
	}
}
func ToMovementResponse(movement entities.Movement) MovementResponse {
	return MovementResponse{
		ID:              movement.ID,
		Number:          movement.Number,
		Type:            movement.Type,
		Total:           movement.Total,
		DateTime:        movement.DateTime.Format("2006-01-02 15:04:05"),
		Description:     movement.Description,
		MovementDetails: ToMovementDetailsResponse(movement.MovementDetail),
	}
}
func ToMovementsResponse(movements []entities.Movement) []MovementResponse {
	var movementsResponse []MovementResponse
	for _, movement := range movements {
		movementsResponse = append(movementsResponse, ToMovementResponse(movement))
	}
	return movementsResponse
}

func ToMovementDetailsResponse(movementDetails []entities.MovementDetail) []MovementDetailResponse {
	var movementDetailsResponse []MovementDetailResponse
	for _, movementDetail := range movementDetails {
		movementDetailsResponse = append(movementDetailsResponse, ToMovementDetailResponse(movementDetail))
	}
	return movementDetailsResponse
}

func ToMovementDetailResponse(movementDetail entities.MovementDetail) MovementDetailResponse {
	var productVariationId *uint
	if movementDetail.ProductVariation != nil {
		productVariationId = &movementDetail.ProductVariation.ID
	} else {
		productVariationId = nil
	}

	var materialId *uint
	if movementDetail.Material != nil {
		materialId = &movementDetail.Material.ID
	} else {
		materialId = nil
	}
	return MovementDetailResponse{
		ID:                 movementDetail.ID,
		ProductVariationID: productVariationId,
		MaterialID:         materialId,
		Quantity:           movementDetail.Quantity,
		Price:              movementDetail.Price,
	}
}
