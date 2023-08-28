package repositories

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
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

func (r *EmployeeRepository) CreateEmployee(employee entities.Employee) error {
	id := r.db.Create(&employee)
	if id.Error != nil {
		return id.Error
	}
	return nil
}

func (r *EmployeeRepository) FindEmployeeByUserId(id uint) (entities.Employee, error) {
	var employee entities.Employee
	res := r.db.Where("user_id = ?", id).Find(&employee)
	if res.Error != nil {
		return entities.Employee{}, res.Error
	}
	return employee, nil
}

func (r *EmployeeRepository) FindAll() ([]entities.Employee, error) {
	var employees []entities.Employee
	r.db.Find(&employees)
	return employees, nil
}

func (r *EmployeeRepository) UpdateEmployee(employee entities.Employee) error {
	r.db.Save(&employee)
	return nil
}

func (r *EmployeeRepository) DeleteEmployee(id string) error {
	r.db.Delete(&entities.Employee{}, id)
	return nil
}
