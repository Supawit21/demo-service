package service

import (
	"Supawit21/demo_service/internal/adapters"
	"Supawit21/demo_service/internal/entity"
)

type EmployeeService struct {
	repo adapters.EmployeeRepository
}

func NewEmployeeService(repo adapters.EmployeeRepository) adapters.EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee *entity.Employee) error {
	return s.repo.CreateEmployee(employee)
}
