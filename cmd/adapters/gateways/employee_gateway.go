package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type EmployeeGateway interface {
	CreateEmployee(employee entities.Employee) error
	FindEmployeeByUserId(id uint) (entities.Employee, error)
	FindAll() ([]entities.Employee, error)
	UpdateEmployee(employee entities.Employee) error
	DeleteEmployee(id string) error
}

type EmployeeGatewayImpl struct {
	repository repositories.EmployeeRepository
}

func (e EmployeeGatewayImpl) CreateEmployee(employee gateway_entities.Employee) error {
	return e.repository.CreateEmployee(employee)
}

func (e EmployeeGatewayImpl) FindEmployeeByUserId(id uint) (entities.Employee, error) {
	employeeDB, err := e.repository.FindEmployeeByUserId(id)

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
		Name: employeeDB.Name,
		DNI:  employeeDB.DNI,
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
	employeesDB, err := e.repository.FindAll()
	if err != nil {
		return nil, err
	}
	employees := make([]entities.Employee, len(employeesDB))
	for i, employeeDB := range employeesDB {
		employees[i] = entities.Employee{
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
