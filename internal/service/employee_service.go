package service

import (
	"Supawit21/demo_service/internal/adapters"
	"Supawit21/demo_service/internal/entity"
	"Supawit21/demo_service/pkg/utils"

	"github.com/google/uuid"
)

type EmployeeService struct {
	repo adapters.EmployeeRepository
}

func NewEmployeeService(repo adapters.EmployeeRepository) adapters.EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee *entity.Employee) error {
	employee.Password, _ = utils.HashPassword(employee.Password)
	return s.repo.CreateEmployee(employee)
}

func (s *EmployeeService) GetEmployee() ([]entity.Employee, error) {
	employee, err := s.repo.GetEmployee()
	return employee, err
}

func (s *EmployeeService) GetEmployeeById(id uuid.UUID) (*entity.Employee, error) {
	employee, err := s.repo.GetEmployeeById(id)
	return employee, err
}

func (s *EmployeeService) UpdateEmployee(id uuid.UUID, employee *entity.Employee) error {
	employee.ID = id
	return s.repo.UpdateEmployee(id, employee)
}
