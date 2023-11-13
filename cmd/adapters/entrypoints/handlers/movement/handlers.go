package movement

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	movement_usecase "github.com/fedeveron01/golang-base/cmd/usecases/movement"
	"github.com/gorilla/mux"
)

type MovementHandlerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllByType(w http.ResponseWriter, r *http.Request)
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

	movements, err := p.movementUseCase.FindAll()
	if err != nil {
		p.WriteErrorResponse(w, err)
	}

	json.NewEncoder(w).Encode(ToMovementsResponse(movements))
}

// GetAll Handle api/movement GET request
func (p *MovementHandler) GetAllByType(w http.ResponseWriter, r *http.Request) {
	var movementRequestByType MovementRequestByType
	err := json.NewDecoder(r.Body).Decode(&movementRequestByType)
	movements, err := p.movementUseCase.FindAllByType(movementRequestByType.IsMaterialMovement)
	if err != nil {
		p.WriteErrorResponse(w, err)
	}

	json.NewEncoder(w).Encode(ToMovementsResponse(movements))
}

// GetById Handle api/movement/{id} GET request
func (p *MovementHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		p.WriteErrorResponse(w, core_errors.NewBadRequestError("id is not valid"))
		return
	}
	movement, err := p.movementUseCase.FindById(uint(uid))
	if err != nil {
		p.WriteErrorResponse(w, err)
	}
	json.NewEncoder(w).Encode(ToMovementResponse(movement))
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
	movement, err = p.movementUseCase.Create(movement, movementRequest.EmployeeID)
	if err != nil {
		p.WriteErrorResponse(w, err)
		return
	}
	movementResponse := ToMovementResponse(movement)
	json.NewEncoder(w).Encode(movementResponse)

}
