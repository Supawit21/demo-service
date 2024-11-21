package repository

import (
	"Supawit21/demo_service/internal/adapters"
	"Supawit21/demo_service/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) adapters.EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) CreateEmployee(employee *entity.Employee) error {
	return r.db.Create(&employee).Error
}

func (r *EmployeeRepository) GetEmployee() ([]entity.Employee, error) {
	var employees []entity.Employee
	result := r.db.Find(&employees)
	return employees, result.Error
}

func (r *EmployeeRepository) GetEmployeeById(id uuid.UUID) (*entity.Employee, error) {
	var employeeById entity.Employee
	result := r.db.First(&employeeById, id)
	return &employeeById, result.Error
}

func (r *EmployeeRepository) UpdateEmployee(id uuid.UUID, employee *entity.Employee) error {
	return r.db.Model(&employee).Updates(employee).Error
}
