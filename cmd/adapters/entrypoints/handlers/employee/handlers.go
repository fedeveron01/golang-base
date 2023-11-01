package employee_handler

import (
	"encoding/json"
	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	employee_usecase "github.com/fedeveron01/golang-base/cmd/usecases/employee"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type EmployeeHandlerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
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
func (e *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !e.IsAuthorized(w, r) {
		return
	}
	employees, err := e.employeeUseCase.FindAll()
	if err != nil {
		e.WriteErrorResponse(w, err)
	}

	employeesResponse := ToEmployeeResponses(employees)
	json.NewEncoder(w).Encode(employeesResponse)
}

// GetById Handle api/employee/{id}
func (e *EmployeeHandler) GetById(w http.ResponseWriter, r *http.Request) {
	if !e.IsAuthorized(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	intId, _ := strconv.ParseInt(id, 10, 64)
	employee, err := e.employeeUseCase.FindById(intId)
	if err != nil {
		e.WriteErrorResponse(w, err)
		return
	}
	employeeResponse := ToEmployeeResponse(employee)
	json.NewEncoder(w).Encode(employeeResponse)
}

// Update Handle api/employee PUT request
func (e *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request) {
	if !e.IsAuthorized(w, r) {
		return
	}
	if !e.IsAdmin(w, r) {
		return
	}
	reqBody, _ := io.ReadAll(r.Body)
	var employeeRequest EmployeeRequest
	json.Unmarshal(reqBody, &employeeRequest)

	employee, err := e.employeeUseCase.UpdateEmployee(ToEmployeeEntity(employeeRequest))
	if err != nil {
		e.WriteErrorResponse(w, err)
		return
	}
	employeeResponse := ToEmployeeResponse(employee)
	json.NewEncoder(w).Encode(employeeResponse)

}
