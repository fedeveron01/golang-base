package charge_handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	charge_usecase "github.com/fedeveron01/golang-base/cmd/usecases/charge"
)

type ChargeHandlerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type ChargeHandler struct {
	entrypoints.HandlerBase
	chargeUseCase charge_usecase.ChargeUseCase
}

func NewChargeHandler(sessionGateway gateways.SessionGateway, chargeUseCase charge_usecase.ChargeUseCase) *ChargeHandler {
	return &ChargeHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		chargeUseCase: chargeUseCase,
	}
}

// Handle api/charge
func (p *ChargeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	charges, err := p.chargeUseCase.FindAll()
	if err != nil {
		p.WriteInternalServerError(w, err)
	}
	chargesResponse := ToChargesResponse(charges)

	json.NewEncoder(w).Encode(chargesResponse)
}

// Handle api/charge
func (p *ChargeHandler) Create(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	if !p.IsAdmin(w, r) {
		return
	}
	reqBody, _ := io.ReadAll(r.Body)
	var charge entities.Charge
	json.Unmarshal(reqBody, &charge)
	err := p.chargeUseCase.CreateCharge(charge)
	if err != nil {
		p.WriteInternalServerError(w, err)
		return
	}

	p.WriteResponse(w, "charge created", http.StatusCreated)
	w.WriteHeader(http.StatusCreated)

}
