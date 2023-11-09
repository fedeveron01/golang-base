package movement

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"time"
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

func ToMovementDetailsResponse(movementDetails []entities.MovementDetail) []MovementDetailResponse {
	var movementDetailsResponse []MovementDetailResponse
	for _, movementDetail := range movementDetails {
		movementDetailsResponse = append(movementDetailsResponse, ToMovementDetailResponse(movementDetail))
	}
	return movementDetailsResponse
}

func ToMovementDetailResponse(movementDetail entities.MovementDetail) MovementDetailResponse {
	return MovementDetailResponse{
		ID:                 movementDetail.ID,
		ProductVariationID: movementDetail.ProductVariation.ID,
		MaterialID:         movementDetail.Material.ID,
		Quantity:           movementDetail.Quantity,
		Price:              movementDetail.Price,
	}
}
