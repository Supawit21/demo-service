package adapters

import "Supawit21/demo_service/internal/entity"

type EmployeeRepository interface {
	CreateEmployee(employee *entity.Employee) error
}
