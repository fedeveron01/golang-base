package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
	"gorm.io/gorm"
)

type EmployeeGatewayImpl struct {
	employeeRepository repositories.EmployeeRepository
}

func NewEmployeeGateway(employeeRepository repositories.EmployeeRepository) *EmployeeGatewayImpl {
	return &EmployeeGatewayImpl{
		employeeRepository: employeeRepository,
	}
}

func (e EmployeeGatewayImpl) CreateEmployee(employee entities.Employee) error {
	employeeDB := gateway_entities.Employee{
		Name:     employee.Name,
		LastName: employee.LastName,
		DNI:      employee.DNI,
		UserId:   employee.User.ID,
		User: gateway_entities.User{
			Model: gorm.Model{
				ID: employee.User.ID,
			},
			UserName: employee.User.UserName,
			Password: employee.User.Password,
		},
		ChargeId: employee.Charge.ID,
		Charge: gateway_entities.Charge{
			Name: employee.Charge.Name,
		},
	}

	return e.employeeRepository.CreateEmployee(employeeDB)
}

func (e EmployeeGatewayImpl) UpdateEmployee(employee entities.Employee) error {

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
	return e.employeeRepository.UpdateEmployee(employeeDB)
}

func (e EmployeeGatewayImpl) FindEmployeeByUserId(id uint) (entities.Employee, error) {
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

	employee := entities.Employee{
		EntitiesBase: core.EntitiesBase{
			ID: employeeDB.ID,
		},
		Name:     employeeDB.Name,
		LastName: employeeDB.LastName,
		DNI:      employeeDB.DNI,
		User: entities.User{
			UserName: employeeDB.User.UserName,
			Password: employeeDB.User.Password,
		},
		Charge: entities.Charge{
			Name: employeeDB.Charge.Name,
		},
		ProductionOrders: productionOrders,
		PurchaseOrders:   purchaseOrders,
	}

	return employee, err
}

func (e EmployeeGatewayImpl) FindAll() ([]entities.Employee, error) {
	employeesDB, err := e.employeeRepository.FindAll()
	if err != nil {
		return nil, err
	}
	employees := make([]entities.Employee, len(employeesDB))
	for i, employeeDB := range employeesDB {
		employees[i] = entities.Employee{
			EntitiesBase: core.EntitiesBase{
				ID: employeeDB.ID,
			},
			Name: employeeDB.Name,
			DNI:  employeeDB.DNI,
			User: entities.User{
				UserName: employeeDB.User.UserName,
				Password: employeeDB.User.Password,
			},
			Charge: entities.Charge{
				Name: employeeDB.Charge.Name,
			},
		}
	}
	return employees, err
}

func (e EmployeeGatewayImpl) DeleteEmployee(id string) error {
	return e.employeeRepository.DeleteEmployee(id)
}
