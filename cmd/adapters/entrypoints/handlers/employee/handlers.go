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

type GetAllEmployeeHandler struct {
	entrypoints.HandlerBase
	employeeUseCase employee_usecase.EmployeeUseCase
}

func NewGetAllEmployeeHandler(sessionGateway gateways.SessionGateway, employeeUseCase employee_usecase.EmployeeUseCase) *GetAllEmployeeHandler {
	return &GetAllEmployeeHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		employeeUseCase: employeeUseCase,
	}
}

// Handle api/employee
func (p *GetAllEmployeeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !p.IsAuthorized(w, r) {
		return
	}
	employees, err := p.employeeUseCase.FindAll()
	if err != nil {
		p.WriteInternalServerError(w, err)
	}

	employeesResponse := ToEmployeeResponses(employees)
	json.NewEncoder(w).Encode(employeesResponse)
}

type GetByIdEmployeeHandler struct {
	entrypoints.HandlerBase
	employeeUseCase employee_usecase.EmployeeUseCase
}

func NewGetByIdEmployeeHandler(sessionGateway gateways.SessionGateway, employeeUseCase employee_usecase.EmployeeUseCase) *GetByIdEmployeeHandler {
	return &GetByIdEmployeeHandler{
		HandlerBase: entrypoints.HandlerBase{
			SessionGateway: sessionGateway,
		},
		employeeUseCase: employeeUseCase,
	}
}

// Handle api/employee/{id}
func (g *GetByIdEmployeeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if !g.IsAuthorized(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 10, 64)
	employee, err := g.employeeUseCase.FindById(intId)
	if err != nil {
		g.WriteInternalServerError(w, err)
		return
	}
	employeeResponse := ToEmployeeResponse(employee)
	json.NewEncoder(w).Encode(employeeResponse)
}
