package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(employee gateway_entities.Employee) error
	FindAll() ([]gateway_entities.Employee, error)
	FindById(id int64) (gateway_entities.Employee, error)
	UpdateEmployee(employee gateway_entities.Employee) (gateway_entities.Employee, error)
	FindEmployeeByUserId(id uint) (gateway_entities.Employee, error)
	DeleteEmployee(id string) error
}

type EmployeeGatewayImpl struct {
	employeeRepository EmployeeRepository
}

func NewEmployeeGateway(employeeRepository EmployeeRepository) *EmployeeGatewayImpl {
	return &EmployeeGatewayImpl{
		employeeRepository: employeeRepository,
	}
}

func (e *EmployeeGatewayImpl) FindAll() ([]entities.Employee, error) {
	employeesDB, err := e.employeeRepository.FindAll()
	if err != nil {
		return nil, err
	}
	employees := make([]entities.Employee, len(employeesDB))
	for i, employeeDB := range employeesDB {
		employees[i] = e.ToBusinessEntity(employeeDB)
	}
	return employees, err
}

func (e *EmployeeGatewayImpl) FindById(id int64) (entities.Employee, error) {
	employeeDB, err := e.employeeRepository.FindById(id)
	if err != nil {
		return entities.Employee{}, err
	}
	employee := e.ToBusinessEntity(employeeDB)
	return employee, err
}

func (e *EmployeeGatewayImpl) CreateEmployee(employee entities.Employee) error {
	employeeDB := e.ToServiceEntity(employee)

	return e.employeeRepository.CreateEmployee(employeeDB)
}

func (e *EmployeeGatewayImpl) UpdateEmployee(employee entities.Employee) (entities.Employee, error) {

	employeeDB := e.ToServiceEntity(employee)
	res, err := e.employeeRepository.UpdateEmployee(employeeDB)
	if err != nil {
		return entities.Employee{}, err
	}
	employee = e.ToBusinessEntity(res)
	return employee, err

}

func (e *EmployeeGatewayImpl) FindEmployeeByUserId(id uint) (entities.Employee, error) {
	employeeDB, err := e.employeeRepository.FindEmployeeByUserId(id)

	employee := e.ToBusinessEntity(employeeDB)

	return employee, err
}

func (e *EmployeeGatewayImpl) DeleteEmployee(id string) error {
	return e.employeeRepository.DeleteEmployee(id)
}

func (e *EmployeeGatewayImpl) ToBusinessEntity(employeeDB gateway_entities.Employee) entities.Employee {
	employee := entities.Employee{
		EntitiesBase: core.EntitiesBase{
			ID: employeeDB.ID,
		},
		Name:     employeeDB.Name,
		LastName: employeeDB.LastName,
		DNI:      employeeDB.DNI,
		User: entities.User{
			EntitiesBase: core.EntitiesBase{
				ID: employeeDB.User.ID,
			},
			UserName: employeeDB.User.UserName,
			Password: employeeDB.User.Password,
			Inactive: employeeDB.User.Inactive,
		},
		Charge: entities.Charge{
			EntitiesBase: core.EntitiesBase{
				ID: employeeDB.Charge.ID,
			},
			Name: employeeDB.Charge.Name,
		},
	}
	return employee
}

func (e *EmployeeGatewayImpl) ToServiceEntity(employee entities.Employee) gateway_entities.Employee {
	employeeDB := gateway_entities.Employee{
		Model: gorm.Model{
			ID: employee.ID,
		},
		Name:     employee.Name,
		LastName: employee.LastName,
		DNI:      employee.DNI,
		UserId:   employee.User.ID,
		ChargeId: employee.Charge.ID,
	}
	return employeeDB
}
