package employee_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
)

type EmployeeUseCase interface {
	FindAll() ([]entities.Employee, error)
	FindById(id int64) (entities.Employee, error)
	CreateEmployee(employee entities.Employee) error
	UpdateEmployee(employee entities.Employee) error
	DeleteEmployee(id string) error
}

type EmployeeGateway interface {
	FindAll() ([]entities.Employee, error)
	CreateEmployee(employee entities.Employee) error
	UpdateEmployee(employee entities.Employee) error
	DeleteEmployee(id string) error
}

type Implementation struct {
	employeeGateway EmployeeGateway
}

func NewEmployeeUsecase(employeeGateway EmployeeGateway) *Implementation {
	return &Implementation{
		employeeGateway: employeeGateway,
	}
}

func (i *Implementation) FindAll() ([]entities.Employee, error) {
	return i.employeeGateway.FindAll()
}
func (i *Implementation) FindById(id int64) (entities.Employee, error) {
	return entities.Employee{}, nil
}
func (i *Implementation) CreateEmployee(employee entities.Employee) error {
	err := i.employeeGateway.CreateEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}
func (i *Implementation) UpdateEmployee(employee entities.Employee) error {
	return i.employeeGateway.UpdateEmployee(employee)
}
func (i *Implementation) DeleteEmployee(id string) error {
	return i.employeeGateway.DeleteEmployee(id)
}
