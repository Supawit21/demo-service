package adapters

import "Supawit21/demo_service/internal/entity"

type EmployeeService interface {
	CreateEmployee(employee *entity.Employee) error
}
