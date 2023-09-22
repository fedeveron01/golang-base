package charge_handler

import (
	"encoding/json"
	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	charge_usecase "github.com/fedeveron01/golang-base/cmd/usecases/charge"
	"io"
	"net/http"
)

type CreateChargeHandler struct {
	entrypoints.HandlerBase
	chargeUseCase charge_usecase.ChargeUseCase
}

type GetAllChargeHandler struct {
	entrypoints.HandlerBase
	chargeUseCase charge_usecase.ChargeUseCase
}

func NewCreateChargeHandler(chargeUseCase charge_usecase.ChargeUseCase) CreateChargeHandler {
	return CreateChargeHandler{
		chargeUseCase: chargeUseCase,
	}
}

// Handle api/charge
func (p CreateChargeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
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

	w.WriteHeader(http.StatusCreated)

}
