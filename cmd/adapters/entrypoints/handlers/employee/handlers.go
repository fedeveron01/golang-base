package employee_handler

import (
	"encoding/json"
	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	employee_usecase "github.com/fedeveron01/golang-base/cmd/usecases/employee"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EmployeeHandlerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type EmployeeHandler struct {
	entrypoints.HandlerBase
	employeeUseCase employee_usecase.EmployeeUseCase
}

func NewEmployeeHandler(sessionGateway gateways.SessionGateway, employeeUseCase employee_usecase.EmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		employeeUseCase: employeeUseCase,
	}
}

// GetAll Handle api/employee
func (p *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	employees, err := p.employeeUseCase.FindAll()
	if err != nil {
		p.WriteErrorResponse(w, err)
	}

	employeesResponse := ToEmployeeResponses(employees)
	json.NewEncoder(w).Encode(employeesResponse)
}

// GetById Handle api/employee/{id}
func (g *EmployeeHandler) GetById(w http.ResponseWriter, r *http.Request) {
	if !g.IsAuthorized(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 10, 64)
	employee, err := g.employeeUseCase.FindById(intId)
	if err != nil {
		g.WriteErrorResponse(w, err)
		return
	}
	employeeResponse := ToEmployeeResponse(employee)
	json.NewEncoder(w).Encode(employeeResponse)
}
