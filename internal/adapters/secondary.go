package adapters

import (
	"Supawit21/demo_service/internal/entity"

	"github.com/google/uuid"
)

type EmployeeRepository interface {
	CreateEmployee(employee *entity.Employee) error
	GetEmployee() ([]entity.Employee, error)
	GetEmployeeById(id uuid.UUID) (*entity.Employee, error)
	UpdateEmployee(id uuid.UUID, employee *entity.Employee) error
}
