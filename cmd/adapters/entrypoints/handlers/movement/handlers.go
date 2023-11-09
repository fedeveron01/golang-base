package movement

import (
	"encoding/json"
	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	movement_usecase "github.com/fedeveron01/golang-base/cmd/usecases/movement"
	"net/http"
)

type MovementHandlerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type MovementHandler struct {
	entrypoints.HandlerBase
	movementUseCase movement_usecase.MovementUseCase
}

func NewMovementHandler(sessionGateway gateways.SessionGateway, movementUseCase movement_usecase.MovementUseCase) *MovementHandler {
	return &MovementHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		movementUseCase: movementUseCase,
	}
}

// GetAll Handle api/movement GET request
func (p *MovementHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	panic("implement me angelo")
}

// GetById Handle api/movement/{id} GET request
func (p *MovementHandler) GetById(w http.ResponseWriter, r *http.Request) {
	panic("implement me angelo")
}

// Create Handle api/movement POST request
func (p *MovementHandler) Create(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	var movementRequest MovementRequest
	err := json.NewDecoder(r.Body).Decode(&movementRequest)
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}
	movement := ToMovement(movementRequest)
	movement, err = p.movementUseCase.Create(movement)
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}
	movementResponse := ToMovementResponse(movement)
	json.NewEncoder(w).Encode(movementResponse)

}
