package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type EmployeeGatewayImpl struct {
	employeeRepository repositories.EmployeeRepository
}

func NewEmployeeGateway(employeeRepository repositories.EmployeeRepository) *EmployeeGatewayImpl {
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
	return entities.Employee{}, nil
}

func (e *EmployeeGatewayImpl) CreateEmployee(employee entities.Employee) error {
	employeeDB := e.ToServiceEntity(employee)

	return e.employeeRepository.CreateEmployee(employeeDB)
}

func (e *EmployeeGatewayImpl) UpdateEmployee(employee entities.Employee) error {

	employeeDB := e.ToServiceEntity(employee)
	return e.employeeRepository.UpdateEmployee(employeeDB)
}

func (e *EmployeeGatewayImpl) FindEmployeeByUserId(id uint) (entities.Employee, error) {
	employeeDB, err := e.employeeRepository.FindEmployeeByUserId(id)

	// map production orders
	productionOrders := make([]entities.ProductionOrder, len(employeeDB.ProductionOrders))
	for i, productionOrderDB := range employeeDB.ProductionOrders {
		productionOrderDetail := make([]entities.ProductionOrderDetail, len(productionOrderDB.ProductionOrderDetail))
		for j, productionOrderDetailDB := range productionOrderDB.ProductionOrderDetail {
			productionOrderDetail[j] = entities.ProductionOrderDetail{
				Quantity: productionOrderDetailDB.Quantity,
				Product: entities.Product{
					Name: productionOrderDetailDB.Product.Name,
				},
			}
		}
		productionOrders[i] = entities.ProductionOrder{
			StartDateTime:         productionOrderDB.StartDateTime,
			EndDateTime:           productionOrderDB.EndDateTime,
			ProductionOrderDetail: productionOrderDetail,
		}

	}

	// map purchase orders
	purchaseOrders := make([]entities.PurchaseOrder, len(employeeDB.PurchaseOrders))
	for i, purchaseOrderDB := range employeeDB.PurchaseOrders {
		purchaseOrderDetail := make([]entities.PurchaseOrderDetail, len(purchaseOrderDB.PurchaseOrderDetails))
		for j, purchaseOrderDetailDB := range purchaseOrderDB.PurchaseOrderDetails {
			purchaseOrderDetail[j] = entities.PurchaseOrderDetail{
				Quantity: purchaseOrderDetailDB.Quantity,
				Material: entities.Material{
					Name: purchaseOrderDetailDB.Material.Name,
				},
			}
		}
		purchaseOrders[i] = entities.PurchaseOrder{
			Number:               purchaseOrderDB.Number,
			PurchaseOrderDetails: purchaseOrderDetail,
		}

	}

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
		Name:     employee.Name,
		LastName: employee.LastName,
		DNI:      employee.DNI,
		User: gateway_entities.User{
			UserName: employee.User.UserName,
			Password: employee.User.Password,
		},
		Charge: gateway_entities.Charge{
			Name: employee.Charge.Name,
		},
	}
	return employeeDB
}
