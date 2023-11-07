package employee_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
)

type EmployeeUseCase interface {
	FindAll() ([]entities.Employee, error)
	FindById(id int64) (entities.Employee, error)
	CreateEmployee(employee entities.Employee) error
	UpdateEmployee(employee entities.Employee) (entities.Employee, error)
	DeleteEmployee(id string) error
}

type EmployeeGateway interface {
	FindAll() ([]entities.Employee, error)
	FindById(id int64) (entities.Employee, error)
	CreateEmployee(employee entities.Employee) error
	UpdateEmployee(employee entities.Employee) (entities.Employee, error)
	DeleteEmployee(id string) error
}

type ChargeGateway interface {
	FindById(id uint) (entities.Charge, error)
}

type Implementation struct {
	employeeGateway EmployeeGateway
	chargeGateway   ChargeGateway
}

func NewEmployeeUseCase(employeeGateway EmployeeGateway, chargeGateway ChargeGateway) *Implementation {
	return &Implementation{
		employeeGateway: employeeGateway,
		chargeGateway:   chargeGateway,
	}
}

func (i *Implementation) FindAll() ([]entities.Employee, error) {
	return i.employeeGateway.FindAll()
}
func (i *Implementation) FindById(id int64) (entities.Employee, error) {
	return i.employeeGateway.FindById(id)
}
func (i *Implementation) CreateEmployee(employee entities.Employee) error {
	err := i.employeeGateway.CreateEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}
func (i *Implementation) UpdateEmployee(employee entities.Employee) (entities.Employee, error) {
	if employee.ID == 0 {
		return entities.Employee{}, core_errors.NewBadRequestError("employee id is not valid")
	}
	if employee.DNI == "" {
		return entities.Employee{}, core_errors.NewBadRequestError("employee dni is not valid")
	}
	if employee.Name == "" {
		return entities.Employee{}, core_errors.NewBadRequestError("employee name is not valid")
	}
	if employee.LastName == "" {
		return entities.Employee{}, core_errors.NewBadRequestError("employee lastname is not valid")
	}

	oldEmployee, err := i.employeeGateway.FindById(int64(employee.ID))
	if err != nil {
		return entities.Employee{}, err
	}
	if oldEmployee.ID == 0 {
		return entities.Employee{}, core_errors.NewNotFoundError("employee not found")
	}

	if employee.Charge.ID != 0 && oldEmployee.Charge.ID != employee.Charge.ID {
		charge, err := i.chargeGateway.FindById(employee.Charge.ID)
		if err != nil {
			return entities.Employee{}, err
		}
		if charge.ID == 0 {
			return entities.Employee{}, core_errors.NewNotFoundError("charge not found")
		}
	} else {
		employee.Charge = oldEmployee.Charge
	}

	employee.User = oldEmployee.User

	return i.employeeGateway.UpdateEmployee(employee)
}
func (i *Implementation) DeleteEmployee(id string) error {
	return i.employeeGateway.DeleteEmployee(id)
}
