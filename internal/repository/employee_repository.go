package repository

import (
	"Supawit21/demo_service/internal/adapters"
	"Supawit21/demo_service/internal/entity"

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
