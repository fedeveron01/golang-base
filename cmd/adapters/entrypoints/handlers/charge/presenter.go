package charge_handler

import "github.com/fedeveron01/golang-base/cmd/core/entities"

func ToChargesResponse(charge []entities.Charge) []ChargeResponse {
	var chargeResponses []ChargeResponse
	for _, charge := range charge {
		chargeResponses = append(chargeResponses, ChargeResponse{
			Id:   float64(charge.ID),
			Name: charge.Name,
		})
	}
	return chargeResponses
}
