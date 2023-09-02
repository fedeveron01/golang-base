package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(database *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: database,
	}
}

func (r *EmployeeRepository) CreateEmployee(employee gateway_entities.Employee) error {
	id := r.db.Create(&employee)
	if id.Error != nil {
		return id.Error
	}
	return nil
}

func (r *EmployeeRepository) FindEmployeeByUserId(id uint) (gateway_entities.Employee, error) {
	var employee gateway_entities.Employee
	res := r.db.Where("user_id = ?", id).Joins("Charge").Find(&employee)
	if res.Error != nil {
		return gateway_entities.Employee{}, res.Error
	}
	return employee, nil
}

func (r *EmployeeRepository) FindAll() ([]gateway_entities.Employee, error) {
	var employees []gateway_entities.Employee
	r.db.Find(&employees)
	return employees, nil
}

func (r *EmployeeRepository) UpdateEmployee(employee gateway_entities.Employee) error {
	r.db.Save(&employee)
	return nil
}

func (r *EmployeeRepository) DeleteEmployee(id string) error {
	r.db.Delete(&gateway_entities.Employee{}, id)
	return nil
}
